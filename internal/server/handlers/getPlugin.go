package handlers

import (
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var GetPlugin = operations.GetPluginHandlerFunc(
	func(params operations.GetPluginParams) middleware.Responder {
		log.V("handlers.GetPlugin()")
		plugin, err := plugin_service.Get(params.PluginName)
		if err != nil {
			return operations.NewGetPluginNotFound().WithPayload(err.Error())
		}
		return operations.NewGetPluginOK().WithPayload(plugin)
	},
)
