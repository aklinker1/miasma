package controllers

import (
	"github.com/aklinker1/miasma/internal/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/services"
	"github.com/go-openapi/runtime/middleware"
)

func UseAppsController(api *operations.MiasmaAPI) {
	// api.GetAppsHandler = getApps
	// api.CreateAppHandler = createApp
	api.GetAppHandler = getApp
	// api.DeleteAppHandler = deleteApp
}

// var getApps = operations.GetAppsHandlerFunc(
// 	func(params operations.GetAppsParams) middleware.Responder {
// 		app, err := services.Files.ReadApp(params.)
// 		if ()
// 		return operations.NewGetAppsOK().WithPayload()
// 	})

// var createApp = operations.CreateAppHandlerFunc(
// 	func(params operations.CreateAppParams) middleware.Responder {
// 		app, err := services.Files.ReadApp()

// 		return *operations.NewCreateAppCreated().WithPayload(app)
// 	})

var getApp = operations.GetAppHandlerFunc(
	func(params operations.GetAppParams) middleware.Responder {
		app, err := services.Files.ReadApp(params.AppName)
		if err != nil {
			return operations.NewGetAppNotFound().WithPayload(err.Error())
		}
		return operations.NewGetAppOK().WithPayload(app)
	})

// var deleteApp = operations.DeleteAppHandlerFunc(
// 	func(params operations.DeleteAppParams) middleware.Responder {
// 		return operations.NewDeleteAppOK().WithPayload()
// 	})
