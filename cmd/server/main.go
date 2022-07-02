package main

import (
	"os"

	"github.com/aklinker1/miasma/internal/server/docker"
	"github.com/aklinker1/miasma/internal/server/fmt"
	"github.com/aklinker1/miasma/internal/server/graphql"
	"github.com/aklinker1/miasma/internal/server/sqlite"
)

// Compile time variables
var (
	VERSION    string
	BUILD      string
	BUILD_HASH string
	BUILD_DATE string
)

func main() {
	logger := &fmt.Logger{}

	db := sqlite.NewDB("/data/miasma/apps.db", logger)
	err := db.Open()
	if err != nil {
		logger.E("Failed to open database: %v", err)
		os.Exit(1)
	}

	runtime, err := docker.NewRuntimeService(logger)
	apps := sqlite.NewAppService(db, runtime, logger)
	env := sqlite.NewEnvService(db, runtime, logger)
	routes := sqlite.NewRouteService(db, logger)
	plugins := sqlite.NewPluginService(db, apps, runtime, logger)
	if err != nil {
		logger.E("Failed to initialize docker runtime: %v", err)
		os.Exit(1)
	}
	resolver := &graphql.Resolver{
		Apps:       apps,
		Routes:     routes,
		EnvService: env,
		Plugins:    plugins,
		Runtime:    runtime,
		Version:    VERSION,
		Logger:     logger,
	}

	server := graphql.NewServer(logger, db, resolver)

	server.ServeGraphql()
}
