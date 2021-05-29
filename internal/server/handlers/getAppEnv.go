package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var GetAppEnv = operations.GetAppEnvHandlerFunc(
	func(params operations.GetAppEnvParams) middleware.Responder {
		log.V("handlers.GetAppEnv()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		appID, err := app_service.GetAppID(db, params.AppName)
		if err != nil {
			return operations.NewGetAppEnvNotFound().WithPayload(err.Error())
		}

		appEnv, err := env_service.Get(db, appID)
		if err != nil {
			return operations.NewGetAppEnvDefault(500).WithPayload(err.Error())
		}

		return operations.NewGetAppEnvOK().WithPayload(appEnv)
	},
)
