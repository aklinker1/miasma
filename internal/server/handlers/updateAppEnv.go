package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/go-openapi/runtime/middleware"
)

var UpdateAppEnv = operations.UpdateAppEnvHandlerFunc(
	func(params operations.UpdateAppEnvParams) middleware.Responder {
		log.V("handlers.UpdateAppEnv()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		newEnv := params.NewEnv.(map[string]interface{})
		if err := validation.AppEnv(newEnv); err != nil {
			return operations.NewUpdateAppEnvBadRequest().
				WithPayload(err.Error())
		}

		appId, err := app_service.GetAppID(tx, params.AppName)
		if err != nil {
			return operations.NewUpdateAppEnvNotFound().WithPayload(err.Error())
		}

		err = env_service.Update(tx, appId, newEnv)
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
		}

		env, err := env_service.Get(tx, appId)
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
		}

		return operations.NewUpdateAppEnvOK().WithPayload(env)
	},
)
