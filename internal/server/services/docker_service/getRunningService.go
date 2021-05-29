package docker_service

import (
	"fmt"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/filters"
	"docker.io/go-docker/api/types/swarm"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func GetRunningService(appName string) (*swarm.Service, error) {
	log.V("docker_service.GetRunningService(%v)", appName)
	filter := filters.NewArgs(
		filters.KeyValuePair{Key: "name", Value: appName},
	)
	services, err := dockerAPI.ServiceList(ctx, types.ServiceListOptions{
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
