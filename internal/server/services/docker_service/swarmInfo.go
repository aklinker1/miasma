package docker_service

import (
	"docker.io/go-docker/api/types/swarm"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func SwarmInfo() *swarm.Swarm {
	log.V("docker_service.SwarmInfo()")
	swarmInfo, err := dockerAPI.SwarmInspect(ctx)
	if err != nil {
		log.E("%v", err)
		return nil
	}
	return &swarmInfo
}
