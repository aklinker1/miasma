package sqlite

import (
	"context"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlitetypes"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/gofrs/uuid"
)

type AppRepo struct {
	Logger server.Logger
}

func (r *AppRepo) GetAll(ctx context.Context, tx server.Tx, filter server.AppsFilter) ([]internal.App, error) {
	var scanned internal.App
	query := sqlb.Select(r.Logger, "apps", map[string]any{
		"id":              &scanned.ID,
		"created_at":      &scanned.CreatedAt,
		"updated_at":      &scanned.UpdatedAt,
		"name":            &scanned.Name,
		"group":           &scanned.Group,
		"system":          &scanned.System,
		"image":           &scanned.Image,
		"image_digest":    &scanned.ImageDigest,
		"auto_upgrade":    &scanned.AutoUpgrade,
		"hidden":          &scanned.Hidden,
		"target_ports":    sqlitetypes.IntArray(&scanned.TargetPorts),
		"published_ports": sqlitetypes.IntArray(&scanned.PublishedPorts),
		"placement":       sqlitetypes.StringArray(&scanned.Placement),
		"volumes":         sqlitetypes.BoundVolumeArray(&scanned.Volumes),
		"command":         sqlitetypes.StringArray(&scanned.Command),
		"networks":        sqlitetypes.StringArray(&scanned.Networks),
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
	if utils.ValueOr(filter.ExcludeHidden, false) {
		query.Where("(hidden = ? OR hidden IS NULL)", 0)
	}
	if filter.AutoUpgrade != nil {
		query.Where("(auto_upgrade = ?)", *filter.AutoUpgrade)
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

func (r *AppRepo) GetOne(ctx context.Context, tx server.Tx, filter server.AppsFilter) (internal.App, error) {
	apps, err := r.GetAll(ctx, tx, filter)
	if err != nil {
		return zero.App, err
	}
	if len(apps) == 0 {
		return zero.App, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "App not found",
			Op:      "sqlite.AppRepo.GetOne",
		}
	}
	return apps[0], nil
}

func (r *AppRepo) Create(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	if app.ID == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return zero.App, err
		}
		app.ID = id.String()
	}
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()

	sql, args := sqlb.Insert(r.Logger, "apps", map[string]any{
		"id":              app.ID,
		"created_at":      app.CreatedAt,
		"updated_at":      app.UpdatedAt,
		"name":            app.Name,
		"group":           app.Group,
		"system":          app.System,
		"image":           app.Image,
		"image_digest":    app.ImageDigest,
		"auto_upgrade":    app.AutoUpgrade,
		"hidden":          app.Hidden,
		"target_ports":    sqlitetypes.IntArray(app.TargetPorts),
		"published_ports": sqlitetypes.IntArray(app.PublishedPorts),
		"placement":       sqlitetypes.StringArray(app.Placement),
		"volumes":         sqlitetypes.BoundVolumeArray(app.Volumes),
		"command":         sqlitetypes.StringArray(app.Command),
		"networks":        sqlitetypes.StringArray(app.Networks),
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return app, err
}

func (r *AppRepo) Update(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	app.UpdatedAt = time.Now()

	sql, args := sqlb.Update(r.Logger, "apps", "id", app.ID, map[string]any{
		"updated_at":      app.UpdatedAt,
		"name":            app.Name,
		"group":           app.Group,
		"system":          app.System,
		"image":           app.Image,
		"image_digest":    app.ImageDigest,
		"auto_upgrade":    app.AutoUpgrade,
		"hidden":          app.Hidden,
		"target_ports":    sqlitetypes.IntArray(app.TargetPorts),
		"published_ports": sqlitetypes.IntArray(app.PublishedPorts),
		"placement":       sqlitetypes.StringArray(app.Placement),
		"volumes":         sqlitetypes.BoundVolumeArray(app.Volumes),
		"command":         sqlitetypes.StringArray(app.Command),
		"networks":        sqlitetypes.StringArray(app.Networks),
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return app, err
}

func (r *AppRepo) Delete(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	app.UpdatedAt = time.Now()
	_, err := tx.ExecContext(ctx, "DELETE FROM apps WHERE id = ?", app.ID)
	return app, err
}
