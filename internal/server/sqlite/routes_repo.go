package sqlite

import (
	"context"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/aklinker1/miasma/internal/server/zero"
)

type RouteRepo struct {
	Logger server.Logger
}

func (r *RouteRepo) GetAll(ctx context.Context, tx server.Tx, filter server.RoutesFilter) ([]internal.Route, error) {
	var scanned internal.Route
	query := sqlb.Select(r.Logger, "routes", map[string]any{
		"app_id":       &scanned.AppID,
		"created_at":   &scanned.CreatedAt,
		"updated_at":   &scanned.UpdatedAt,
		"host":         &scanned.Host,
		"path":         &scanned.Path,
		"traefik_rule": &scanned.TraefikRule,
	})
	if filter.AppID != nil {
		query.Where("app_id = ?", *filter.AppID)
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, server.NewDatabaseError("findRoutes", err)
	}
	dest := query.ScanDest()
	result := make([]internal.Route, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, server.NewDatabaseError("findRoutes", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func (r *RouteRepo) GetOne(ctx context.Context, tx server.Tx, filter server.RoutesFilter) (internal.Route, error) {
	routes, err := r.GetAll(ctx, tx, filter)
	if err != nil {
		return zero.Route, err
	}
	if len(routes) == 0 {
		return zero.Route, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "Route not found",
			Op:      "sqlite.RouteRepo.GetOne",
		}
	}
	return routes[0], nil
}

func (r *RouteRepo) Create(ctx context.Context, tx server.Tx, route internal.Route) (internal.Route, error) {
	route.CreatedAt = time.Now()
	route.UpdatedAt = time.Now()

	sql, args := sqlb.Insert(r.Logger, "routes", map[string]any{
		"app_id":       route.AppID,
		"created_at":   route.CreatedAt,
		"updated_at":   route.UpdatedAt,
		"host":         route.Host,
		"path":         route.Path,
		"traefik_rule": route.TraefikRule,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return route, err
}

func (r *RouteRepo) Update(ctx context.Context, tx server.Tx, route internal.Route) (internal.Route, error) {
	route.UpdatedAt = time.Now()

	sql, args := sqlb.Update(r.Logger, "routes", "app_id", route.AppID, map[string]any{
		"created_at":   route.CreatedAt,
		"updated_at":   route.UpdatedAt,
		"host":         route.Host,
		"path":         route.Path,
		"traefik_rule": route.TraefikRule,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return route, err
}

func (r *RouteRepo) Delete(ctx context.Context, tx server.Tx, route internal.Route) (internal.Route, error) {
	route.UpdatedAt = time.Now()
	_, err := tx.ExecContext(ctx, "DELETE FROM routes WHERE app_id = ?", route.AppID)
	return route, err
}
