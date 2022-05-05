package graphql

import "github.com/aklinker1/miasma/internal/server"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Apps    server.AppService
	Runtime server.RuntimeService
}