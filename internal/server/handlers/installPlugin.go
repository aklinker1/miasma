package handlers

import (
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var InstallPlugin = operations.InstallPluginHandlerFunc(
	func(params operations.InstallPluginParams) middleware.Responder {
		log.V("handlers.InstallPlugin()")
		plugin, err := plugin_service.Install(params.PluginName)
		if err != nil {
			return operations.NewInstallPluginDefault(500).WithPayload(err.Error())
		}
		return operations.NewInstallPluginCreated().WithPayload(plugin)
	},
)
