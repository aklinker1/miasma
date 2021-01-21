package services

import (
	"context"
	"fmt"

	goDocker "docker.io/go-docker"
	goSwarm "docker.io/go-docker/api/types/swarm"
)

type dockerService struct{}

var Docker = &dockerService{}

var docker *goDocker.Client
var swarm *goDocker.SwarmAPIClient
var ctx = context.Background()

func init() {
	var err error
	docker, err = goDocker.NewEnvClient()
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

func (service *dockerService) SwarmInfo() *goSwarm.Swarm {
	swarmInfo, err := docker.SwarmInspect(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &swarmInfo
}
