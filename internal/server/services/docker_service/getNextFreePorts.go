package docker_service

import (
	"fmt"

	"docker.io/go-docker/api/types"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func GetNextFreePorts(count int) ([]uint32, error) {
	log.V("docker_service.GetNextFreePorts(%v)", count)
	services, err := dockerAPI.ServiceList(ctx, types.ServiceListOptions{})
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
