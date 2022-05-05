package graphql

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	defaultPort = 3000
)

type graphqlServer struct {
	logger   server.Logger
	port     int
	db       server.DB
	resolver gqlgen.ResolverRoot
}

func NewServer(logger server.Logger, db server.DB, resolver gqlgen.ResolverRoot) server.Server {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.I("$PORT could be not converted to integer ('%s'), falling back to %d", portStr, defaultPort)
		port = defaultPort
	}
	return &graphqlServer{
		logger:   logger,
		port:     port,
		db:       db,
		resolver: resolver,
	}
}

// ServeGraphql implements server.Server
func (s *graphqlServer) ServeGraphql() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/graphql", s.createGraphqlHandler())
	r.Handle("/playground", playground.Handler("Anime Skip", "/graphql"))

	s.logger.I("Miasma server started at :%d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
}

func (s *graphqlServer) createGraphqlHandler() *handler.Server {
	// Define the server
	es := gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers: s.resolver,
	})
	srv := handler.NewDefaultServer(es)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})

	// Enable introspection
	srv.Use(extension.Introspection{})

	// Setup error handling
	srv.SetRecoverFunc(func(ctx context.Context, e interface{}) error {
		var message string
		var err error
		var ok bool
		if err, ok = e.(error); ok {
			message = err.Error()
			s.logger.E("Paniced: %v", err)
			debug.PrintStack()
		} else {
			message = fmt.Sprintf("Unhandled error %v (%T)", e, e)
		}
		return &server.Error{
			Code:    server.EINTERNAL,
			Message: message,
			Err:     err,
		}
	})

	return srv
}
