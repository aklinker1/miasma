package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/server/services/run_config_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/go-openapi/runtime/middleware"
)

var UpdateRunConfig = operations.UpdateRunConfigHandlerFunc(
	func(params operations.UpdateRunConfigParams) middleware.Responder {
		log.V("handlers.UpdateRunConfig()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		if err := validation.RunConfig(params.NewRunConfig); err != nil {
			return operations.NewUpdateAppEnvBadRequest().
				WithPayload(err.Error())
		}

		appID, err := app_service.GetAppID(tx, params.AppName)
		if err != nil {
			return operations.NewUpdateRunConfigNotFound().WithPayload(err.Error())
		}

		newRunConfig, err := run_config_service.Get(tx, appID)
		newRunConfig.Command = params.NewRunConfig.Command
		newRunConfig.Networks = params.NewRunConfig.Networks
		newRunConfig.Placement = params.NewRunConfig.Placement
		newRunConfig.PublishedPorts = params.NewRunConfig.PublishedPorts
		newRunConfig.TargetPorts = params.NewRunConfig.TargetPorts
		newRunConfig.Volumes = params.NewRunConfig.Volumes
		err = run_config_service.Upsert(tx, newRunConfig)
		if err != nil {
			return operations.NewUpdateRunConfigDefault(500).WithPayload(err.Error())
		}

		details, err := app_service.Details(tx, params.AppName)
		if err != nil {
			return operations.NewUpdateRunConfigNotFound().WithPayload(err.Error())
		}

		env, err := env_service.Get(tx, appID)
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
		}

		plugins, err := plugin_service.GetAppPlugins(tx, appID)
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
		}

		err = app_service.Reload(details, env, plugins)
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
		}

		return operations.NewUpdateRunConfigOK().WithPayload(details.RunConfig)
	},
)
