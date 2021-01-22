package services

import (
	"context"
	"fmt"

	dockerLib "docker.io/go-docker"
	dockerTypes "docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/filters"
	dockerSwarmTypes "docker.io/go-docker/api/types/swarm"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/log"
	"github.com/aklinker1/miasma/internal/server/utils/mappers"
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

func (service *dockerService) StartApp(app *models.App) error {
	existingService, _ := service.GetRunningService(*app.Name)
	if existingService != nil {
		return fmt.Errorf("%s is already started", *app.Name)
	}
	newService := mappers.App.ToService(app)
	options := dockerTypes.ServiceCreateOptions{
		QueryRegistry: true,
	}
	_, err := docker.ServiceCreate(ctx, *newService, options)
	if err != nil {
		return err
	}
	return nil
}

func (service *dockerService) StopApp(app *models.App) error {
	runningService, err := service.GetRunningService(*app.Name)
	if err != nil {
		return err
	}
	return docker.ServiceRemove(ctx, runningService.ID)
}
