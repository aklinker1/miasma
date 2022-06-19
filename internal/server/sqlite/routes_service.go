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
	db server.DB
}

func NewRouteService(db server.DB) server.RouteService {
	return &RouteService{
		db: db,
	}
}

// FindRoute implements server.RouteService
func (s *RouteService) FindRoute(ctx context.Context, filter server.RoutesFilter) (internal.AppRouting, error) {
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyRoute, err
	}
	return findRoute(ctx, tx, filter)
}
