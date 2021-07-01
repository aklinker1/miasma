package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var StopApp = operations.StopAppHandlerFunc(
	func(params operations.StopAppParams) middleware.Responder {
		log.V("handlers.StopApp()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		app, _ := app_service.Get(db, params.AppName)
		if app == nil {
			return operations.NewStopAppNotFound().WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}

		err := app_service.Stop(app)
		if err != nil {
			return operations.NewStopAppDefault(500).WithPayload(err.Error())
		}
		return operations.NewStopAppNoContent()
	},
)
