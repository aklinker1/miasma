package handlers

import (
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var ListPlugins = operations.ListPluginsHandlerFunc(
	func(params operations.ListPluginsParams) middleware.Responder {
		log.V("handlers.ListPlugins()")
		plugins, err := plugin_service.List()
		if err != nil {
			return operations.NewInstallPluginDefault(500).WithPayload(err.Error())
		}
		return operations.NewListPluginsOK().WithPayload(plugins)
	},
)
