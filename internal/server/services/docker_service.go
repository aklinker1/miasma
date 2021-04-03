package services

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	dockerLib "docker.io/go-docker"
	dockerTypes "docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/filters"
	dockerSwarmTypes "docker.io/go-docker/api/types/swarm"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/mappers"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared/log"
)

type dockerService struct{}

var Docker = &dockerService{}

var docker *dockerLib.Client
var ctx = context.Background()

func init() {
	var err error
	docker, err = dockerLib.NewEnvClient()
	if err != nil {
		panic("Could not connect to host's docker service")
	}
}

func (service *dockerService) Version() *string {
	version, err := docker.ServerVersion(ctx)
	if err != nil {
		log.E("%v", err)
		return nil
	}
	versionString := fmt.Sprintf("%s-%s", version.Version, version.GitCommit)
	return &versionString
}

func (service *dockerService) SwarmInfo() *dockerSwarmTypes.Swarm {
	swarmInfo, err := docker.SwarmInspect(ctx)
	if err != nil {
		log.E("%v", err)
		return nil
	}
	return &swarmInfo
}

func (service *dockerService) GetRunningService(appName string) (*dockerSwarmTypes.Service, error) {
	filter := filters.NewArgs(
		filters.KeyValuePair{Key: "name", Value: appName},
	)
	services, err := docker.ServiceList(ctx, dockerTypes.ServiceListOptions{
		Filters: filter,
	})
	if err != nil {
		return nil, err
	}
	for _, s := range services {
		if s.Spec.Annotations.Name == appName {
			return &s, nil
		}
	}
	return nil, fmt.Errorf("%s is not running", appName)
}

func (service *dockerService) IsAppServiceRunning(appName string) bool {
	runningService, _ := service.GetRunningService(appName)
	return runningService != nil
}

func (service *dockerService) CreateNetworkIfNotAvailable(networkName string) error {
	networks, err := docker.NetworkList(ctx, dockerTypes.NetworkListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key: "name", Value: networkName,
		}),
	})
	if err != nil {
		return err
	}
	if len(networks) > 0 {
		log.V("Network '%s' already exists, not creating", networkName)
		return nil
	}
	_, err = docker.NetworkCreate(ctx, networkName, dockerTypes.NetworkCreate{
		Driver: "overlay",
		Scope:  "swarm",
		Labels: map[string]string{
			"miasma": "true",
		},
	})
	if err != nil {
		return err
	}
	log.V("Created network: %s", networkName)
	return nil
}

func (service *dockerService) DestroyNetwork(appName string) error {
	return docker.NetworkRemove(ctx, appName)
}

func (service *dockerService) StartApp(app *types.AppMetaData) error {
	existingService, _ := service.GetRunningService(app.Name)
	if existingService != nil {
		return fmt.Errorf("%s is already running", app.Name)
	}
	pluginMeta, err := Plugin.GetPluginMeta()
	if err != nil {
		return err
	}
	digest, err := Docker.GetDigest(app.Image)
	if err != nil {
		return err
	}
	newService, err := mappers.App.ToService(app, pluginMeta, App.getNextPorts(app), digest)
	if err != nil {
		return err
	}
	options := dockerTypes.ServiceCreateOptions{
		QueryRegistry: true,
	}

	return service.StartService(*newService, options)
}

func (service *dockerService) StartService(serviceSpec dockerSwarmTypes.ServiceSpec, options dockerTypes.ServiceCreateOptions) error {
	_, err := docker.ServiceCreate(ctx, serviceSpec, options)
	return err
}

func (service *dockerService) StopService(serviceName string) error {
	runningService, err := service.GetRunningService(serviceName)
	if err != nil {
		return err
	}
	return docker.ServiceRemove(ctx, runningService.ID)
}

func (service *dockerService) StopApp(app *models.App) error {
	return service.StopService(*app.Name)
}

func (service *dockerService) UpdateService(existingService *dockerSwarmTypes.Service, newServiceSpec *dockerSwarmTypes.ServiceSpec) error {
	_, err := docker.ServiceUpdate(
		ctx,
		existingService.ID,
		existingService.Version,
		*newServiceSpec,
		dockerTypes.ServiceUpdateOptions{
			QueryRegistry: false,
		},
	)
	if err != nil {
		log.E("Failed to update service: %v", err)
		return err
	}

	return nil
}

func (service *dockerService) GetNextAvailablePorts(count int) ([]uint32, error) {
	services, err := docker.ServiceList(ctx, dockerTypes.ServiceListOptions{})
	if err != nil {
		return nil, err
	}
	filledPorts := map[uint32]bool{}
	for _, service := range services {
		for _, port := range service.Endpoint.Ports {
			filledPorts[port.PublishedPort] = true
		}
	}
	results := []uint32{}
	var port uint32
	for port = 3001; port < 4000 && len(results) < count; port++ {
		if _, ok := filledPorts[port]; !ok {
			results = append(results, port)
		}
	}
	if len(results) < count {
		return nil, fmt.Errorf("Not enough available ports to start the service (required=%d, available=%d)", count, len(results))
	}
	return results, nil
}

func (service *dockerService) PullImage(baseImage string) error {
	log.V("Pulling %s...", baseImage)
	ctx := context.Background()

	stream, err := dockerLib.ImageAPIClient.ImagePull(docker, ctx, baseImage, dockerTypes.ImagePullOptions{})
	if err != nil {
		log.E("Failed to pull image: %v", err)
		return err
	}
	defer stream.Close()

	// Read 1 image at a time (if pulling more than 1)
	// reader := bufio.NewReader(stream)
	// for ok := true; ok; {
	// 	line, _, err := reader.ReadLine()
	// 	log.V("%s", string(line))
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			log.E("Failed to read line: %v", err)
	// 		}
	// 		ok = false
	// 	}
	// }
	// Do it all
	_, err = ioutil.ReadAll(stream)
	if err != nil {
		log.E("Failed to save pulled image: %v", err)
		return err
	}
	log.V("Done!")
	return nil
}

func (service *dockerService) GetDigest(baseImage string) (string, error) {
	log.V("Inspecting %s...", baseImage)
	ctx := context.Background()

	info, _, err := dockerLib.ImageAPIClient.ImageInspectWithRaw(docker, ctx, baseImage)
	if err != nil {
		log.E("Failed to inspect image: %v", err)
		return "", err
	}
	for _, digest := range info.RepoDigests {
		if strings.Contains(digest, "@sha256:") {
			digest := digest[strings.LastIndex(digest, "@")+1:]
			log.V("Digest: %v", digest)
			return digest, nil
		}
	}
	return "", errors.New("Could not find digest with hash instead of tag")
}
