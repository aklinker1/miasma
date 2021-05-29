package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var StartApp = operations.StartAppHandlerFunc(
	func(params operations.StartAppParams) middleware.Responder {
		log.V("handlers.StartApp()")
		var err error
		db, onDefer := database.ReadOnly(&err)
		defer onDefer()

		details, err := app_service.Details(db, params.AppName)
		if err != nil {
			return operations.NewStartAppNotFound().WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}
		appId := details.App.ID.String()

		plugins, err := plugin_service.GetAppPlugins(db, appId)
		if err != nil {
			return operations.NewStartAppDefault(500).WithPayload(err.Error())
		}

		err = app_service.Start(db, details, plugins)
		if err != nil {
			return operations.NewStartAppDefault(500).WithPayload(err.Error())
		}
		return operations.NewStartAppNoContent()
	},
)
