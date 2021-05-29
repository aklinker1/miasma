package server

import (
	"flag"

	"github.com/go-openapi/loads"

	"github.com/aklinker1/miasma/internal/server/gen/restapi"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/handlers"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func Start() {
	// Load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal("%v", err)
	}

	// Create new service API
	api := operations.NewMiasmaAPI(swaggerSpec)
	api.Logger = log.I
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Parse flags
	var portFlag = flag.Int("port", 3000, "Port to run this service on")
	flag.Parse()
	server.Port = *portFlag

	// Use handlers
	useControllers(api)
	server.ConfigureAPI()

	// Serve API
	if err := server.Serve(); err != nil {
		log.Fatal("%v", err)
	}
}

func useControllers(api *operations.MiasmaAPI) {
	// Health
	api.HealthCheckHandler = handlers.HealthCheck

	// Apps
	api.ListAppsHandler = handlers.ListApps
	api.CreateAppHandler = handlers.CreateApp
	api.GetAppHandler = handlers.GetApp
	api.UpgradeAppHandler = handlers.UpgradeApp
	api.DeleteAppHandler = handlers.DeleteApp
	api.StartAppHandler = handlers.StartApp
	api.StopAppHandler = handlers.StopApp
	api.EditAppHandler = handlers.EditApp

	// App Env
	api.GetAppEnvHandler = handlers.GetAppEnv
	api.UpdateAppEnvHandler = handlers.UpdateAppEnv

	// Run Config
	api.GetRunConfigHandler = handlers.GetRunConfig
	api.UpdateRunConfigHandler = handlers.UpdateRunConfig

	// Plugins
	api.ListPluginsHandler = handlers.ListPlugins
	api.GetPluginHandler = handlers.GetPlugin
	api.InstallPluginHandler = handlers.InstallPlugin
	api.UninstallPluginHandler = handlers.UninstallPlugin
	// Traefik
	api.GetAppTraefikConfigHandler = handlers.GetAppTraefikConfig
	api.UpdateAppTraefikConfigHandler = handlers.UpdateAppTraefikConfig
	api.RemoveAppTraefikConfigHandler = handlers.RemoveAppTraefikConfig
}
