package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/aklinker1/miasma/package/models"
	"github.com/go-openapi/runtime/middleware"
)

var EditApp = operations.EditAppHandlerFunc(
	func(params operations.EditAppParams) middleware.Responder {
		log.V("handlers.EditApp()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		app, err := app_service.Get(tx, params.AppName)
		if err != nil {
			return operations.NewEditAppNotFound().WithPayload(err.Error())
		}

		err = validation.AppEdit(params.NewApp)
		if err != nil {
			return operations.NewEditAppBadRequest().WithPayload(err.Error())
		}

		newApp := &models.App{
			ID:     app.ID,
			Image:  app.Image,
			Name:   params.NewApp.Name,
			Group:  params.NewApp.Group,
			Hidden: params.NewApp.Hidden,
		}
		err = app_service.Upsert(tx, newApp)
		if err != nil {
			return operations.NewEditAppDefault(500).WithPayload(err.Error())
		}

		return operations.NewEditAppOK().WithPayload(newApp)
	},
)
