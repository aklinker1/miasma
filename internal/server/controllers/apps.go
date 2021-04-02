package controllers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/go-openapi/runtime/middleware"
)

func UseAppsController(api *operations.MiasmaAPI) {
	api.GetAppsHandler = getApps
	api.CreateAppHandler = createApp
	api.GetAppHandler = getApp
	api.DeleteAppHandler = deleteApp
	api.StartAppHandler = startApp
	api.StopAppHandler = stopApp
	api.GetAppConfigHandler = getAppConfig
	api.UpdateAppConfigHandler = updateAppConfig
	api.GetAppEnvHandler = getAppEnv
	api.UpdateAppEnvHandler = updateAppEnv
	api.UpdateAppHandler = updateApp
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
		err := validation.AppName(*inputApp.Name)
		if err != nil {
			return operations.NewCreateAppBadRequest().WithPayload(err.Error())
		}
		existingApp, _ := services.App.Get(*inputApp.Name)
		if existingApp != nil {
			return operations.NewCreateAppBadRequest().WithPayload(fmt.Sprintf("%s already exists", *inputApp.Name))
		}

		err = services.Docker.PullImage(*inputApp.Image)
		if err != nil {
			return operations.NewCreateAppBadRequest().WithPayload(err.Error())
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
		app, _ := services.App.GetAppMeta(params.AppName)
		if app == nil {
			return operations.NewStartAppNotFound().WithPayload(fmt.Sprintf("%s does not exist", params.AppName))
		}

		err := services.Docker.StartApp(app)
		if err != nil {
			return operations.NewStartAppDefault(500).WithPayload(err.Error())
		}
		return operations.NewStartAppOK()
	})

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
	})

var getAppConfig = operations.GetAppConfigHandlerFunc(
	func(params operations.GetAppConfigParams) middleware.Responder {
		appConfig, err := services.App.GetConfig(params.AppName)
		if err != nil {
			return operations.NewGetAppConfigNotFound().WithPayload(err.Error())
		}
		return operations.NewGetAppConfigOK().WithPayload(appConfig)
	})

var updateAppConfig = operations.UpdateAppConfigHandlerFunc(
	func(params operations.UpdateAppConfigParams) middleware.Responder {
		_, err := services.App.GetConfig(params.AppName)
		if err != nil {
			return operations.NewUpdateAppConfigNotFound().WithPayload(err.Error())
		}
		err = validation.AppConfig(params.NewAppConfig)
		if err != nil {
			return operations.NewUpdateAppConfigBadRequest().WithPayload(err.Error())
		}
		appConfig, err := services.App.UpdateConfig(params.AppName, params.NewAppConfig)
		if err != nil {
			return operations.NewUpdateAppConfigDefault(500).WithPayload(err.Error())
		}
		return operations.NewUpdateAppConfigOK().WithPayload(appConfig)
	})

var getAppEnv = operations.GetAppEnvHandlerFunc(
	func(params operations.GetAppEnvParams) middleware.Responder {
		appEnv, err := services.App.GetEnv(params.AppName)
		if err != nil {
			return operations.NewGetAppEnvNotFound().WithPayload(err.Error())
		}
		return operations.NewGetAppEnvOK().WithPayload(appEnv)
	})

var updateAppEnv = operations.UpdateAppEnvHandlerFunc(
	func(params operations.UpdateAppEnvParams) middleware.Responder {
		_, err := services.App.GetConfig(params.AppName)
		if err != nil {
			return operations.NewUpdateAppEnvNotFound().WithPayload(err.Error())
		}
		err = validation.AppEnv(params.NewEnv.(map[string]interface{}))
		if err != nil {
			return operations.NewUpdateAppEnvBadRequest().WithPayload(err.Error())
		}
		env, err := services.App.UpdateEnv(params.AppName, params.NewEnv.(map[string]interface{}))
		if err != nil {
			return operations.NewUpdateAppEnvDefault(500).WithPayload(err.Error())
		}
		return operations.NewUpdateAppEnvOK().WithPayload(env)
	})

var updateApp = operations.UpdateAppHandlerFunc(
	func(params operations.UpdateAppParams) middleware.Responder {
		appMeta, err := services.App.GetAppMeta(params.AppName)
		if err != nil {
			return operations.NewUpdateAppBadRequest().WithPayload(err.Error())
		}

		var newImage string
		if params.NewImage == nil {
			newImage = appMeta.Image
		} else {
			newImage = *params.NewImage
		}
		updated, err := services.App.UpdateAndReload(appMeta, newImage)

		if err != nil {
			return operations.NewUpdateAppDefault(500).WithPayload(err.Error())
		}
		if !updated {
			return operations.NewUpdateAppBadRequest().WithPayload(fmt.Sprintf("No updates are available for %s!", newImage))
		}
		return operations.NewUpdateAppOK().WithPayload(nil)
	})
