package main

import (
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/fmt"
	"github.com/aklinker1/miasma/internal/server/graphql"
)

func main() {
	logger := &fmt.Logger{}
	var db server.DB
	resolver := &graphql.Resolver{}

	server := graphql.NewServer(logger, db, resolver)

	server.ServeGraphql()
}
