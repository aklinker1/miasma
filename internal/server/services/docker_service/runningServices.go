package docker_service

import (
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/swarm"
)

func RunningServices() ([]swarm.Service, error) {
	services, err := dockerAPI.ServiceList(ctx, types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}
