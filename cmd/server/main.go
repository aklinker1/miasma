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

// Environment Variables
var (
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
)

// Other constants
var (
	dataDir          = "/data/miasma"
	databasePath     = dataDir + "/apps.db"
	certResolverName = "miasmaresolver"
)

func main() {
	logger := &fmt.Logger{}

	db := sqlite.NewDB(databasePath, logger)
	err := db.Open()
	if err != nil {
		logger.E("Failed to open database: %v", err)
		os.Exit(1)
	}

	runtime, err := docker.NewRuntimeService(logger, certResolverName)
	apps := sqlite.NewAppService(db, runtime, logger)
	env := sqlite.NewEnvService(db, runtime, logger)
	routes := sqlite.NewRouteService(db, logger)
	plugins := sqlite.NewPluginService(db, apps, runtime, logger, dataDir, certResolverName)
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

	server := graphql.NewServer(logger, db, resolver, ACCESS_TOKEN)

	server.ServeGraphql()
}
