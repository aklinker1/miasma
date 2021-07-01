package docker_service

import (
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/swarm"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func StartService(serviceSpec swarm.ServiceSpec, options types.ServiceCreateOptions) error {
	log.V("docker_service.StartService(%v, options:***)", serviceSpec)

	_, err := dockerAPI.ServiceCreate(ctx, serviceSpec, options)
	return err
}
