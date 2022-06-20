package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

var (
	EmptyRoute = internal.AppRouting{}
)

type RouteService struct {
	db     server.DB
	logger server.Logger
}

func NewRouteService(db server.DB, logger server.Logger) server.RouteService {
	return &RouteService{
		db:     db,
		logger: logger,
	}
}

// FindRoute implements server.RouteService
func (s *RouteService) FindRoute(ctx context.Context, filter server.RoutesFilter) (internal.AppRouting, error) {
	s.logger.D("Finding route that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyRoute, err
	}
	return findRoute(ctx, tx, filter)
}
