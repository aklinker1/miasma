package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/run_config_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var GetRunConfig = operations.GetRunConfigHandlerFunc(
	func(params operations.GetRunConfigParams) middleware.Responder {
		log.V("handlers.GetRunConfig()")
		var err error
		db, onDefer := database.ReadOnly(&err)
		defer onDefer()

		appID, err := app_service.GetAppID(db, params.AppName)
		if err != nil {
			return operations.NewGetRunConfigNotFound().WithPayload(err.Error())
		}

		appConfig, err := run_config_service.Get(db, appID)
		if err != nil {
			return operations.NewGetRunConfigNotFound().WithPayload(err.Error())
		}

		return operations.NewGetRunConfigOK().WithPayload(appConfig)
	},
)
