package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/internal/shared/constants"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"github.com/go-openapi/runtime/middleware"
)

var HealthCheck = operations.HealthCheckHandlerFunc(
	func(params operations.HealthCheckParams) middleware.Responder {
		log.V("handlers.HealthCheck()")
		dockerVersion := docker_service.Version()
		swarm := docker_service.SwarmInfo()

		return operations.NewHealthCheckOK().WithPayload(&models.Health{
			Version:       shared.StringPtr(fmt.Sprintf("v%v-%v-%v", constants.VERSION, constants.BUILD_DATE, constants.BUILD_HASH)),
			DockerVersion: dockerVersion,
			Swarm: &models.HealthSwarm{
				ID:          swarm.ID,
				CreatedAt:   swarm.CreatedAt.String(),
				UpdatedAt:   swarm.UpdatedAt.String(),
				JoinCommand: fmt.Sprintf("docker swarm join --token %s <main-node-ip:port>", swarm.JoinTokens.Worker),
			},
		})
	},
)
