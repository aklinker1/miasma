package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var UninstallPlugin = operations.UninstallPluginHandlerFunc(
	func(params operations.UninstallPluginParams) middleware.Responder {
		log.V("handlers.UninstallPlugin()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		plugin, pluginApp, err := plugin_service.Uninstall(tx, params.PluginName)
		if err != nil {
			return operations.NewUninstallPluginDefault(500).WithPayload(err.Error())
		}

		err = app_service.Stop(pluginApp)
		if err != nil {
			return operations.NewUninstallPluginDefault(500).WithPayload(err.Error())
		}

		return operations.NewUninstallPluginOK().WithPayload(plugin)
	},
)
