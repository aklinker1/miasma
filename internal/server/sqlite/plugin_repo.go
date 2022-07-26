package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlitetypes"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/samber/lo"
)

type PluginRepo struct {
	Logger server.Logger
}

func (r *PluginRepo) GetAll(ctx context.Context, tx server.Tx, filter server.PluginsFilter) ([]internal.Plugin, error) {
	var scanned internal.Plugin
	query := sqlb.Select(r.Logger, "plugins", map[string]any{
		"name":    sqlitetypes.PluginName(&scanned.Name),
		"enabled": &scanned.Enabled,
		"config":  sqlitetypes.JSON(&scanned.Config),
	})
	if filter.Name != nil {
		query.Where("name = ?", *filter.Name)
	}
	if filter.NameContains != nil {
		query.Where("name ILIKE ?", "%"+*filter.NameContains+"%")
	}
	if filter.Enabled != nil {
		query.Where("enabled = ?", lo.Ternary(*filter.Enabled, 1, 0))
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, server.NewDatabaseError("findPlugins", err)
	}
	dest := query.ScanDest()
	result := make([]internal.Plugin, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, server.NewDatabaseError("findPlugins", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func (r *PluginRepo) GetOne(ctx context.Context, tx server.Tx, filter server.PluginsFilter) (internal.Plugin, error) {
	plugins, err := r.GetAll(ctx, tx, filter)
	if err != nil {
		return zero.Plugin, err
	}
	if len(plugins) == 0 {
		return zero.Plugin, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "Plugin not found",
			Op:      "sqlite.PluginRepo.GetOne",
		}
	}
	return plugins[0], nil
}

func (r *PluginRepo) GetTraefik(ctx context.Context, tx server.Tx) (internal.Plugin, error) {
	return r.GetOne(ctx, tx, server.PluginsFilter{
		Name: lo.ToPtr(internal.PluginNameTraefik),
	})
}

func (r *PluginRepo) Update(ctx context.Context, tx server.Tx, plugin internal.Plugin) (internal.Plugin, error) {
	data := map[string]any{
		"enabled": plugin.Enabled,
		"config":  sqlitetypes.JSON(plugin.Config),
	}
	sql, args := sqlb.Update(r.Logger, "plugins", "name", sqlitetypes.PluginName(plugin.Name), data).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return plugin, err
}
