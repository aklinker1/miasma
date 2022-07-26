package graphql

import "github.com/aklinker1/miasma/internal/server"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	DB                 server.DB
	AppRepo            server.AppRepo
	PluginRepo         server.PluginRepo
	EnvRepo            server.EnvRepo
	RouteRepo          server.RouteRepo
	RuntimeRepo        server.RuntimeRepo
	RuntimeServiceRepo server.RuntimeServiceRepo
	RuntimeNodeRepo    server.RuntimeNodeRepo
	RuntimeTaskRepo    server.RuntimeTaskRepo
	RuntimeImageRepo   server.RuntimeImageRepo
	Logger             server.Logger
	Version            string
	CertResolverName   string
}
