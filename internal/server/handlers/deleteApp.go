package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var DeleteApp = operations.DeleteAppHandlerFunc(
	func(params operations.DeleteAppParams) middleware.Responder {
		log.V("handlers.DeleteApp()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		app, _ := app_service.Get(tx, params.AppName)
		if app == nil {
			return operations.NewDeleteAppNotFound().
				WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}

		err = app_service.Delete(tx, app)
		if err != nil {
			return operations.NewDeleteAppDefault(500).WithPayload(err.Error())
		}

		return operations.NewDeleteAppOK().WithPayload(app)
	},
)
