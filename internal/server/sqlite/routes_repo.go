package sqlite

import (
	"context"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/samber/lo"
)

func findRoutes(ctx context.Context, tx server.Tx, filter server.RoutesFilter) ([]internal.Route, error) {
	var scanned internal.Route
	query := sqlb.Select("routes", map[string]any{
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

func findRoute(ctx context.Context, tx server.Tx, filter server.RoutesFilter) (internal.Route, error) {
	routes, err := findRoutes(ctx, tx, filter)
	if err != nil {
		return EmptyRoute, err
	}
	if len(routes) == 0 {
		return EmptyRoute, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "Route not found",
			Op:      "findRoute",
		}
	}
	return routes[0], nil
}

func findRouteOrNil(ctx context.Context, tx server.Tx, filter server.RoutesFilter) (*internal.Route, error) {
	routes, err := findRoutes(ctx, tx, filter)
	if err != nil {
		return nil, err
	} else if len(routes) == 0 {
		return nil, nil
	}
	return lo.ToPtr(routes[0]), nil
}

func createRoute(ctx context.Context, tx server.Tx, route internal.Route) (internal.Route, error) {
	route.CreatedAt = time.Now()
	route.UpdatedAt = time.Now()

	sql, args := sqlb.Insert("routes", map[string]any{
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

func updateRoute(ctx context.Context, tx server.Tx, route internal.Route) (internal.Route, error) {
	route.UpdatedAt = time.Now()

	sql, args := sqlb.Update("routes", "app_id", route.AppID, map[string]any{
		"created_at":   route.CreatedAt,
		"updated_at":   route.UpdatedAt,
		"host":         route.Host,
		"path":         route.Path,
		"traefik_rule": route.TraefikRule,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return route, err
}

func deleteRoute(ctx context.Context, tx server.Tx, route internal.Route) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM routes WHERE app_id = ?", route.AppID)
	return err
}
