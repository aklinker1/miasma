package docker_service

import (
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/swarm"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func UpdateService(existingService *swarm.Service, newServiceSpec *swarm.ServiceSpec) error {
	log.V("docker_service.UpdateService(%v, %v)", existingService, newServiceSpec)
	_, err := dockerAPI.ServiceUpdate(
		ctx,
		existingService.ID,
		existingService.Version,
		*newServiceSpec,
		types.ServiceUpdateOptions{
			QueryRegistry: false,
		},
	)
	if err != nil {
		log.E("Failed to update service: %v", err)
		return err
	}

	return nil
}
