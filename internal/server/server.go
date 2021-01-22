package server

import (
	"flag"

	"github.com/go-openapi/loads"

	"github.com/aklinker1/miasma/internal/server/controllers"
	"github.com/aklinker1/miasma/internal/server/gen/restapi"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/utils/log"
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
	controllers.UseHealthController(api)
	controllers.UseAppsController(api)
}
