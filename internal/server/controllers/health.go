package controllers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/go-openapi/runtime/middleware"
)

func UseHealthController(api *operations.MiasmaAPI) {
	api.GetHealthCheckHandler = getHealthCheck
}

var getHealthCheck = operations.GetHealthCheckHandlerFunc(
	func(params operations.GetHealthCheckParams) middleware.Responder {
		dockerVersion := services.Docker.Version()
		swarm := services.Docker.SwarmInfo()

		return operations.NewGetHealthCheckOK().WithPayload(&models.Health{
			Version:       &constants.VERSION,
			DockerVersion: dockerVersion,
			Swarm: &models.HealthSwarm{
				ID:          swarm.ID,
				CreatedAt:   swarm.CreatedAt.String(),
				UpdatedAt:   swarm.UpdatedAt.String(),
				JoinCommand: fmt.Sprintf("docker swarm join --token %s <miasma-ip:port>", swarm.JoinTokens.Worker),
			},
		})
	})
