package controllers

import (
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/go-openapi/runtime/middleware"
)

func UsePluginsController(api *operations.MiasmaAPI) {
	api.ListPluginsHandler = listPlugins
	api.GetPluginHandler = getPlugin
	api.InstallPluginHandler = installPlugin
	api.UninstallPluginHandler = uninstallPlugin
}

var listPlugins = operations.ListPluginsHandlerFunc(
	func(params operations.ListPluginsParams) middleware.Responder {
		plugins, err := services.Plugin.ListAll()
		if err != nil {
			return operations.NewInstallPluginDefault(500).WithPayload(err.Error())
		}
		return operations.NewListPluginsOK().WithPayload(plugins)
	})

var getPlugin = operations.GetPluginHandlerFunc(
	func(params operations.GetPluginParams) middleware.Responder {
		pluginMeta, err := services.Plugin.GetPluginMeta()
		if err != nil {
			return operations.NewGetPluginDefault(404).WithPayload(err.Error())
		}
		plugin, err := services.Plugin.Get(params.PluginName, pluginMeta)
		if err != nil {
			return operations.NewGetPluginNotFound().WithPayload(err.Error())
		}
		return operations.NewGetPluginOK().WithPayload(plugin)
	})

var installPlugin = operations.InstallPluginHandlerFunc(
	func(params operations.InstallPluginParams) middleware.Responder {
		plugin, err := services.Plugin.Install(params.PluginName)
		if err != nil {
			return operations.NewInstallPluginDefault(500).WithPayload(err.Error())
		}
		return operations.NewInstallPluginCreated().WithPayload(plugin)
	})

var uninstallPlugin = operations.UninstallPluginHandlerFunc(
	func(params operations.UninstallPluginParams) middleware.Responder {
		plugin, err := services.Plugin.Uninstall(params.PluginName)
		if err != nil {
			return operations.NewUninstallPluginDefault(500).WithPayload(err.Error())
		}
		return operations.NewUninstallPluginOK().WithPayload(plugin)
	})
