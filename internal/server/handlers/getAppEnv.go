package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var GetAppEnv = operations.GetAppEnvHandlerFunc(
	func(params operations.GetAppEnvParams) middleware.Responder {
		log.V("handlers.GetAppEnv()")
		var err error
		db, onDefer := database.ReadOnly(&err)
		defer onDefer()

		appEnv, err := env_service.Get(db, params.AppName)
		if err != nil {
			return operations.NewGetAppEnvNotFound().WithPayload(err.Error())
		}

		return operations.NewGetAppEnvOK().WithPayload(appEnv)
	},
)
