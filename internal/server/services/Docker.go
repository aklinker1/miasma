package services

import (
	"context"
	"fmt"

	dockerLib "docker.io/go-docker"
	dockerTypes "docker.io/go-docker/api/types"
	dockerSwarmTypes "docker.io/go-docker/api/types/swarm"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/mappers"
)

type dockerService struct{}

var Docker = &dockerService{}

var docker *dockerLib.Client
var swarm *dockerLib.SwarmAPIClient
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
		fmt.Println(err)
		return nil
	}
	versionString := fmt.Sprintf("%s-%s", version.Version, version.GitCommit)
	return &versionString
}

func (service *dockerService) SwarmInfo() *dockerSwarmTypes.Swarm {
	swarmInfo, err := docker.SwarmInspect(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &swarmInfo
}

func (service *dockerService) StartAppName(appName string) error {
	app, err := App.Get(appName)
	if err != nil {
		return err
	}
	return service.StartApp(app)
}

func (service *dockerService) StartApp(app *models.App) error {
	newService := mappers.App.ToService(app)
	options := dockerTypes.ServiceCreateOptions{
		QueryRegistry: true,
	}
	createdService, err := docker.ServiceCreate(ctx, *newService, options)
	if err != nil {
		return err
	}
	fmt.Println(createdService)
	return nil
}
