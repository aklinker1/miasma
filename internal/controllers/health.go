package controllers

import (
	"github.com/aklinker1/miasma/internal/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/utils/constants"
	"github.com/go-openapi/runtime/middleware"
)

func UseHealthController(api *operations.MiasmaAPI) {
	api.GetHealthCheckHandler = getHealthCheck
}

var getHealthCheck = operations.GetHealthCheckHandlerFunc(
	func(params operations.GetHealthCheckParams) middleware.Responder {
		return operations.NewGetHealthCheckOK().WithPayload(&operations.GetHealthCheckOKBody{
			Version: &constants.VERSION,
		})
	})
