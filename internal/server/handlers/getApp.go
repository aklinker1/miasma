package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var GetApp = operations.GetAppHandlerFunc(
	func(params operations.GetAppParams) middleware.Responder {
		log.V("handlers.GetApp()")
		var err error
		db, onDefer := database.ReadOnly(&err)
		defer onDefer()

		app, err := app_service.Get(db, params.AppName)
		if err != nil {
			return operations.NewGetAppNotFound().WithPayload(err.Error())
		}
		return operations.NewGetAppOK().WithPayload(app)
	},
)
