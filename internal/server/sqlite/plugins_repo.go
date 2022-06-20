package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlb"
	"github.com/aklinker1/miasma/internal/server/sqlite/sqlitetypes"
	"github.com/samber/lo"
)

func findPlugins(ctx context.Context, tx server.Tx, filter server.PluginsFilter) ([]internal.Plugin, error) {
	var scanned internal.Plugin
	query := sqlb.Select("plugins", map[string]any{
		"name":    sqlitetypes.PluginName(&scanned.Name),
		"enabled": &scanned.Enabled,
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

func findPlugin(ctx context.Context, tx server.Tx, filter server.PluginsFilter) (internal.Plugin, error) {
	plugins, err := findPlugins(ctx, tx, filter)
	if err != nil {
		return EmptyPlugin, err
	}
	if len(plugins) == 0 {
		return EmptyPlugin, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "Plugin not found",
			Op:      "findPlugin",
		}
	}
	return plugins[0], nil
}

func updatePlugin(ctx context.Context, tx server.Tx, plugin internal.Plugin) (internal.Plugin, error) {
	sql, args := sqlb.Update("plugins", "name", sqlitetypes.PluginName(plugin.Name), map[string]any{
		"enabled": plugin.Enabled,
	}).ToSQL()
	_, err := tx.ExecContext(ctx, sql, args...)
	return plugin, err
}
