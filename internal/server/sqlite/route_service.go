package sqlite

import (
	"context"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

var (
	EmptyRoute = internal.Route{}
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
func (s *RouteService) FindRoute(ctx context.Context, filter server.RoutesFilter) (internal.Route, error) {
	s.logger.D("Finding route that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyRoute, err
	}
	return findRoute(ctx, tx, filter)
}

// FindRoute implements server.RouteService
func (s *RouteService) FindRouteOrNil(ctx context.Context, filter server.RoutesFilter) (*internal.Route, error) {
	s.logger.D("Finding route (or falling back to nil) that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return nil, err
	}
	return findRouteOrNil(ctx, tx, filter)
}

// SetRoute implements server.RouteService
func (s *RouteService) Create(ctx context.Context, route internal.Route) (internal.Route, error) {
	s.logger.D("Creating route: %s", route.AppID)
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyRoute, err
	}
	defer tx.Rollback()

	route.CreatedAt = time.Now()
	route.UpdatedAt = time.Now()
	created, err := createRoute(ctx, tx, route)
	if err != nil {
		return EmptyRoute, err
	}

	tx.Commit()
	return created, err
}

// SetRoute implements server.RouteService
func (s *RouteService) Update(ctx context.Context, route internal.Route) (internal.Route, error) {
	s.logger.D("Updating route: %s", route.AppID)
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyRoute, err
	}
	defer tx.Rollback()

	route.UpdatedAt = time.Now()
	created, err := updateRoute(ctx, tx, route)
	if err != nil {
		return EmptyRoute, err
	}

	tx.Commit()
	return created, err
}

// DeleteRoute implements server.RouteService
func (s *RouteService) Delete(ctx context.Context, route internal.Route) (internal.Route, error) {
	s.logger.D("Deleting route: %s", route.AppID)
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyRoute, err
	}
	defer tx.Rollback()

	err = deleteRoute(ctx, tx, route)
	if err != nil {
		return EmptyRoute, err
	}

	tx.Commit()
	return route, err
}
