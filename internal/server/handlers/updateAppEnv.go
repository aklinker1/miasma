package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
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

		details, err := app_service.Details(tx, params.AppName)
		if err != nil {
			return operations.NewUpdateAppEnvNotFound().WithPayload(err.Error())
		}
		appID := details.App.ID.String()

		err = env_service.Update(tx, appID, newEnv)
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
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

		return operations.NewUpdateAppEnvOK().WithPayload(env)
	},
)
