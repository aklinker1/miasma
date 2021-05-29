package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var UpgradeApp = operations.UpgradeAppHandlerFunc(
	func(params operations.UpgradeAppParams) middleware.Responder {
		log.V("handlers.UpgradeApp()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		details, err := app_service.Details(tx, params.AppName)
		if err != nil {
			return operations.NewUpgradeAppNotFound().WithPayload(err.Error())
		}
		appId := details.App.ID.String()
		env, err := env_service.Get(tx, appId)
		if err != nil {
			return operations.NewUpgradeAppDefault(500).WithPayload(err.Error())
		}
		plugins, err := plugin_service.GetAppPlugins(tx, appId)
		if err != nil {
			return operations.NewUpgradeAppDefault(500).WithPayload(err.Error())
		}

		var newImage string
		if params.NewImage == nil {
			newImage = details.App.Image
		} else {
			newImage = *params.NewImage
		}
		updated, err := app_service.Upgrade(details, env, plugins, newImage)

		if err != nil {
			return operations.NewUpgradeAppDefault(500).WithPayload(err.Error())
		}
		if !updated {
			return operations.NewUpgradeAppBadRequest().WithPayload(fmt.Sprintf("No updates are available for %s!", newImage))
		}
		return operations.NewUpgradeAppOK().WithPayload(nil)
	},
)
