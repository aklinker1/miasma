package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var InstallPlugin = operations.InstallPluginHandlerFunc(
	func(params operations.InstallPluginParams) middleware.Responder {
		log.V("handlers.InstallPlugin()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		plugin, pluginDetails, err := plugin_service.Install(tx, params.PluginName)
		if err != nil {
			return operations.NewInstallPluginDefault(500).WithPayload(err.Error())
		}

		plugins, err := plugin_service.GetAppPlugins(tx, pluginDetails.App.ID.String())
		if err != nil {
			return operations.NewStartAppDefault(500).WithPayload(err.Error())
		}

		err = app_service.Start(tx, pluginDetails, plugins)
		if err != nil {
			return operations.NewStartAppDefault(500).WithPayload(err.Error())
		}

		return operations.NewInstallPluginCreated().WithPayload(plugin)
	},
)
