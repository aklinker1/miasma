package controllers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/go-openapi/runtime/middleware"
)

func UseAppsController(api *operations.MiasmaAPI) {
	api.GetAppsHandler = getApps
	api.CreateAppHandler = createApp
	api.GetAppHandler = getApp
	api.DeleteAppHandler = deleteApp
	api.StartAppHandler = startApp
	api.StopAppHandler = stopApp
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

var createApp = operations.CreateAppHandlerFunc(
	func(params operations.CreateAppParams) middleware.Responder {
		if params.App == nil {
			return operations.NewCreateAppBadRequest().WithPayload("Request body is required")
		}
		inputApp := *params.App
		existingApp, _ := services.App.Get(*inputApp.Name)
		if existingApp != nil {
			return operations.NewCreateAppBadRequest().WithPayload(fmt.Sprintf("%s already exists", *inputApp.Name))
		}

		newApp, err := services.App.Create(inputApp)
		if err != nil {
			return operations.NewCreateAppDefault(500).WithPayload(err.Error())
		}

		return operations.NewCreateAppCreated().WithPayload(newApp)
	})

var getApp = operations.GetAppHandlerFunc(
	func(params operations.GetAppParams) middleware.Responder {
		app, err := services.App.Get(params.AppName)
		if err != nil {
			return operations.NewGetAppNotFound().WithPayload(err.Error())
		}
		return operations.NewGetAppOK().WithPayload(app)
	})

var deleteApp = operations.DeleteAppHandlerFunc(
	func(params operations.DeleteAppParams) middleware.Responder {
		app, _ := services.App.Get(params.AppName)
		if app == nil {
			return operations.NewDeleteAppNotFound().WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}

		err := services.App.Delete(app)
		if err != nil {
			return operations.NewDeleteAppDefault(500).WithPayload(err.Error())
		}

		return operations.NewDeleteAppOK().WithPayload(app)
	})

var startApp = operations.StartAppHandlerFunc(
	func(params operations.StartAppParams) middleware.Responder {
		app, _ := services.App.Get(params.AppName)
		if app == nil {
			return operations.NewStartAppNotFound().WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}

		err := services.Docker.StartApp(app)
		if err != nil {
			return operations.NewStartAppDefault(500).WithPayload(err.Error())
		}
		return operations.NewStartAppOK()
	},
)

var stopApp = operations.StopAppHandlerFunc(
	func(params operations.StopAppParams) middleware.Responder {
		app, _ := services.App.Get(params.AppName)
		if app == nil {
			return operations.NewStopAppNotFound().WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}

		err := services.Docker.StopApp(app)
		if err != nil {
			return operations.NewStopAppDefault(500).WithPayload(err.Error())
		}
		return operations.NewStopAppOK()
	},
)
