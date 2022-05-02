package main

import (
	"os"

	"github.com/aklinker1/miasma/internal/server"
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

	resolver := &graphql.Resolver{}

	server := graphql.NewServer(logger, db, resolver)

	server.ServeGraphql()
}
