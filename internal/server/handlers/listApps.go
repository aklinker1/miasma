package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var ListApps = operations.ListAppsHandlerFunc(
	func(params operations.ListAppsParams) middleware.Responder {
		log.V("handlers.ListApps()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		showHidden := params.Hidden != nil && *params.Hidden
		apps, err := app_service.List(db, showHidden)
		if err != nil {
			return operations.NewListAppsDefault(500).WithPayload(err.Error())
		}

		return operations.NewListAppsOK().WithPayload(apps)
	},
)
