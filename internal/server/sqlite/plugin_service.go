package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

var (
	EmptyPlugin = internal.Plugin{}
)

type PluginService struct {
	db      server.DB
	logger  server.Logger
	apps    server.AppService
	runtime server.RuntimeService
}

func NewPluginService(db server.DB, apps server.AppService, runtime server.RuntimeService, logger server.Logger) server.PluginService {
	return &PluginService{
		db:      db,
		logger:  logger,
		apps:    apps,
		runtime: runtime,
	}
}

func (s *PluginService) setEnabled(ctx context.Context, plugin internal.Plugin, enabled bool) (internal.Plugin, error) {
	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyPlugin, err
	}
	defer tx.Rollback()

	plugin.Enabled = enabled
	updated, err := updatePlugin(ctx, tx, plugin)
	if err != nil {
		return EmptyPlugin, err
	}

	// Restart managed apps
	apps, err := findApps(ctx, tx, server.AppsFilter{})
	if err != nil {
		return EmptyPlugin, err
	}
	s.runtime.RestartRunningApps(ctx, apps)
	if err != nil {
		return EmptyPlugin, err
	}

	tx.Commit()
	return updated, nil
}

// DisablePlugin implements server.PluginService
func (s *PluginService) DisablePlugin(ctx context.Context, plugin internal.Plugin) (internal.Plugin, error) {
	return s.setEnabled(ctx, plugin, false)
}

// EnablePlugin implements server.PluginService
func (s *PluginService) EnablePlugin(ctx context.Context, plugin internal.Plugin) (internal.Plugin, error) {
	return s.setEnabled(ctx, plugin, true)
}

// FindPlugins implements server.PluginService
func (s *PluginService) FindPlugins(ctx context.Context, filter server.PluginsFilter) ([]internal.Plugin, error) {
	s.logger.D("Finding plugins that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return nil, err
	}
	return findPlugins(ctx, tx, filter)
}

// FindPlugin implements server.PluginService
func (s *PluginService) FindPlugin(ctx context.Context, filter server.PluginsFilter) (internal.Plugin, error) {
	s.logger.D("Finding plugins that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyPlugin, err
	}
	return findPlugin(ctx, tx, filter)
}
