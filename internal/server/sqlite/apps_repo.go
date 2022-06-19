package sqlite

import (
	"context"
	"fmt"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlb"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/gofrs/uuid"
)

func findApps(ctx context.Context, tx server.Tx, filter server.AppsFilter) ([]internal.App, error) {
	var scanned internal.App
	query := sqlb.Select("apps", map[string]any{
		"id":              &scanned.ID,
		"created_at":      &scanned.CreatedAt,
		"updated_at":      &scanned.UpdatedAt,
		"name":            &scanned.Name,
		"\"group\"":       &scanned.Group,
		"image":           &scanned.Image,
		"image_digest":    &scanned.ImageDigest,
		"hidden":          &scanned.Hidden,
		"target_ports":    &scanned.TargetPorts,
		"published_ports": &scanned.PublishedPorts,
		"placement":       &scanned.Placement,
		"volumes":         &scanned.Volumes,
		"networks":        &scanned.Networks,
		"routing":         &scanned.Routing,
		"command":         &scanned.Command,
	})
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.Name != nil {
		query.Where("name = ?", *filter.Name)
	}
	if filter.NameContains != nil {
		query.Where("name ILIKE ?", "%"+*filter.NameContains+"%")
	}
	if !utils.BoolOr(filter.IncludeHidden, false) {
		query.Where("(hidden = ? OR hidden IS NULL)", 0)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	if filter.Sort != nil {
		query.OrderBy(*filter.Sort)
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, server.NewDatabaseError("findApps", err)
	}
	dest := query.ScanDest()
	result := make([]internal.App, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, server.NewDatabaseError("findApps", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findApp(ctx context.Context, tx server.Tx, filter server.AppsFilter) (internal.App, error) {
	apps, err := findApps(ctx, tx, filter)
	if err != nil {
		return EmptyApp, err
	}
	if len(apps) == 0 {
		return EmptyApp, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "App not found",
			Op:      "findApp",
		}
	}
	return apps[0], nil
}

func createApp(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return EmptyApp, err
	}
	app.ID = id.String()
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	fmt.Println(app.Routing)

	sql, args := sqlb.Insert("apps", map[string]any{
		"id":              app.ID,
		"created_at":      app.CreatedAt,
		"updated_at":      app.UpdatedAt,
		"name":            app.Name,
		"\"group\"":       app.Group,
		"image":           app.Image,
		"image_digest":    app.ImageDigest,
		"hidden":          app.Hidden,
		"target_ports":    utils.ToJSON(app.TargetPorts),
		"published_ports": utils.ToJSON(app.PublishedPorts),
		"placement":       utils.ToJSON(app.Placement),
		"volumes":         utils.ToJSON(app.Volumes),
		"networks":        utils.ToJSON(app.Networks),
		"routing":         utils.ToJSON(app.Routing),
		"command":         app.Command,
	}).ToSQL()
	_, err = tx.ExecContext(ctx, sql, args...)
	return app, err
}

func updateApp(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	app.UpdatedAt = time.Now()

	sql, args := sqlb.Update("apps", app.ID, map[string]any{
		"updated_at":      app.UpdatedAt,
		"name":            app.Name,
		"\"group\"":       app.Group,
		"image":           app.Image,
		"image_digest":    app.ImageDigest,
		"hidden":          app.Hidden,
		"target_ports":    utils.ToJSON(app.TargetPorts),
		"published_ports": utils.ToJSON(app.PublishedPorts),
		"placement":       utils.ToJSON(app.Placement),
		"volumes":         utils.ToJSON(app.Volumes),
		"networks":        utils.ToJSON(app.Networks),
		"routing":         utils.ToJSON(app.Routing),
		"command":         app.Command,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return app, err
}
