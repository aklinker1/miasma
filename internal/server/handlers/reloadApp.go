package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var ReloadApp = operations.ReloadAppHandlerFunc(
	func(params operations.ReloadAppParams) middleware.Responder {
		log.V("handlers.ReloadApp()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		app, err := app_service.Details(db, params.AppName)
		if err != nil {
			return operations.NewReloadAppNotFound().WithPayload(err.Error())
		}
		appID := app.App.ID.String()
		plugins, err := plugin_service.GetAppPlugins(db, appID)
		if err != nil {
			return operations.NewReloadAppDefault(500).WithPayload(err.Error())
		}

		if docker_service.IsServiceRunning(params.AppName) {
			env, err := env_service.Get(db, appID)
			if err != nil {
				return operations.NewReloadAppDefault(500).WithPayload(err.Error())
			}
			err = app_service.Reload(app, env, plugins)
			if err != nil {
				return operations.NewReloadAppDefault(500).WithPayload(err.Error())
			}
		} else {
			err = app_service.Start(db, app, plugins)
			if err != nil {
				return operations.NewReloadAppDefault(500).WithPayload(err.Error())
			}
		}

		return operations.NewReloadAppCreated()

	},
)
