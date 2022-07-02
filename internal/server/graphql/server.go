package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/web"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const (
	defaultPort = 3000
)

type graphqlServer struct {
	logger      server.Logger
	port        int
	db          server.DB
	resolver    gqlgen.ResolverRoot
	accessToken string
}

func NewServer(logger server.Logger, db server.DB, resolver gqlgen.ResolverRoot, accessToken string) server.Server {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.I("$PORT could be not converted to integer ('%s'), falling back to %d", portStr, defaultPort)
		port = defaultPort
	}
	return &graphqlServer{
		logger:      logger,
		port:        port,
		db:          db,
		resolver:    resolver,
		accessToken: accessToken,
	}
}

// ServeGraphql implements server.Server
func (s *graphqlServer) ServeGraphql() error {
	requireCredentials := s.accessToken != ""

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
	}))
	if requireCredentials {
		r.Use(s.tokenAuthMiddleware)
	} else {
		s.logger.W("Server running without credentials")
	}

	r.Handle("/graphql", s.createGraphqlHandler())
	r.Handle("/playground", playground.Handler("Miasma", "/graphql"))
	r.Get("/*", web.Handler("/"))

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

func (s *graphqlServer) tokenAuthMiddleware(next http.Handler) http.Handler {
	s.logger.I("Server initialized with simple token based credentials")
	headerPrefix := "Bearer "

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeGraphqlError(
				rw,
				"Authorization header required, but was not passed. See https://aklinker1.github.io/miasma/authorization for more details",
				http.StatusUnauthorized,
			)
			return
		}
		if !strings.HasPrefix(authHeader, headerPrefix) {
			writeGraphqlError(
				rw,
				"Authorization header format incorrect, must be \"Bearer <token>\". See https://aklinker1.github.io/miasma/authorization for more details",
				http.StatusUnauthorized,
			)
			return
		}

		token := strings.Replace(authHeader, headerPrefix, "", 1)
		if token != s.accessToken {
			writeGraphqlError(rw, "Authorization token is not valid", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func writeJson(rw http.ResponseWriter, data any, status int) {
	rw.Header().Add("Content-Type", "application/json")
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	rw.WriteHeader(status)
	rw.Write(body)
}

func writeGraphqlError(rw http.ResponseWriter, message string, status int) {
	writeJson(rw, map[string]any{
		"errors": []map[string]any{{
			"message": message,
		}},
	}, status)
}
