package main

import (
	"os"

	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/docker"
	"github.com/aklinker1/miasma/internal/server/fmt"
	"github.com/aklinker1/miasma/internal/server/graphql"
	"github.com/aklinker1/miasma/internal/server/sqlite"
)

func main() {
	logger := &fmt.Logger{}

	db := sqlite.NewDB("/data/miasma/apps.db", logger)
	err := db.Open()
	if err != nil {
		logger.E("Failed to open database: %v", server.ExternalErrorMessage(err))
		os.Exit(1)
	}

	apps := sqlite.NewAppService(db)
	runtime, err := docker.NewRuntimeService()
	if err != nil {
		logger.E("Failed to initialize docker runtime: %v", server.ExternalErrorMessage(err))
		os.Exit(1)
	}
	resolver := &graphql.Resolver{
		Apps:    apps,
		Runtime: runtime,
	}

	server := graphql.NewServer(logger, db, resolver)

	server.ServeGraphql()
}
