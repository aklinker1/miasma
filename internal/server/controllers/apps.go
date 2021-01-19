package controllers

import (
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/go-openapi/runtime/middleware"
)

func UseAppsController(api *operations.MiasmaAPI) {
	api.GetAppsHandler = getApps
	// api.CreateAppHandler = createApp
	api.GetAppHandler = getApp
	// api.DeleteAppHandler = deleteApp
}

var getApps = operations.GetAppsHandlerFunc(
	func(params operations.GetAppsParams) middleware.Responder {
		showHidden := params.Hidden != nil && *params.Hidden
		apps, err := services.App.GetAll(showHidden)
		if err != nil {
			return operations.NewGetAppsDefault(500).WithPayload(err.Error())
		}
		return operations.NewGetAppsOK().WithPayload(apps)
	})

// var createApp = operations.CreateAppHandlerFunc(
// 	func(params operations.CreateAppParams) middleware.Responder {
// 		_, err := services.Apps.Get()

// 		return *operations.NewCreateAppCreated().WithPayload(app)
// 	})

var getApp = operations.GetAppHandlerFunc(
	func(params operations.GetAppParams) middleware.Responder {
		app, err := services.App.Get(params.AppName)
		if err != nil {
			return operations.NewGetAppNotFound().WithPayload(err.Error())
		}
		return operations.NewGetAppOK().WithPayload(app)
	})

// var deleteApp = operations.DeleteAppHandlerFunc(
// 	func(params operations.DeleteAppParams) middleware.Responder {
// 		return operations.NewDeleteAppOK().WithPayload()
// 	})
