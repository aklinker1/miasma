package docker_service

import (
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/filters"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func UpsertNetwork(networkName string) error {
	log.V("docker_service.UpsertNetwork(%v)", networkName)
	networks, err := dockerAPI.NetworkList(ctx, types.NetworkListOptions{
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
	_, err = dockerAPI.NetworkCreate(ctx, networkName, types.NetworkCreate{
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
