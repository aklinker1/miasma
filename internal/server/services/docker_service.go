package services

import (
	"context"
	"fmt"

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

func (service *dockerService) CreateNetwork(appName string) error {
	_, err := docker.NetworkCreate(ctx, appName, dockerTypes.NetworkCreate{
		Driver: "overlay",
		Scope:  "swarm",
	})
	if err != nil {
		return err
	}
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
	newService, err := mappers.App.ToService(app, service.GetNextAvailablePorts)
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

func (service *dockerService) UpdateService(serviceName string, newServiceSpec *dockerSwarmTypes.ServiceSpec) error {
	existingService, err := service.GetRunningService(serviceName)
	if err != nil {
		return err
	}
	_, err = docker.ServiceUpdate(
		ctx,
		existingService.ID,
		existingService.Version,
		*newServiceSpec,
		dockerTypes.ServiceUpdateOptions{
			QueryRegistry: false,
		},
	)
	if err != nil {
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
