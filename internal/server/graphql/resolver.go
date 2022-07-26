package graphql

import (
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/services"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Version string
	DB      server.DB
	Logger  server.Logger

	AppRepo            server.AppRepo
	PluginRepo         server.PluginRepo
	EnvRepo            server.EnvRepo
	RouteRepo          server.RouteRepo
	RuntimeRepo        server.RuntimeRepo
	RuntimeServiceRepo server.RuntimeServiceRepo
	RuntimeNodeRepo    server.RuntimeNodeRepo
	RuntimeTaskRepo    server.RuntimeTaskRepo
	RuntimeImageRepo   server.RuntimeImageRepo

	AppService     *services.AppService
	PluginService  *services.PluginService
	RuntimeService *services.RuntimeService
}
