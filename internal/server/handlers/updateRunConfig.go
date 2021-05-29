package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
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

		appId, err := app_service.GetAppID(tx, params.AppName)
		if err != nil {
			return operations.NewUpdateRunConfigNotFound().WithPayload(err.Error())
		}

		err = run_config_service.Upsert(tx, params.NewRunConfig)
		if err != nil {
			return operations.NewUpdateRunConfigDefault(500).WithPayload(err.Error())
		}

		runConfig, err := run_config_service.Get(tx, appId)
		if err != nil {
			return operations.NewUpdateRunConfigDefault(500).WithPayload(err.Error())
		}

		return operations.NewUpdateRunConfigOK().WithPayload(runConfig)
	},
)
