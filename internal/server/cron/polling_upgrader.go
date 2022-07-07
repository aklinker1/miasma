package cron

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/samber/lo"
)

type PollingUpgrader struct {
	Logger  server.Logger
	Apps    server.AppService
	Runtime server.RuntimeService
	Routes  server.RouteService
	Plugins server.PluginService
	Env     server.EnvService
}

func (u *PollingUpgrader) Cron() {
	u.Logger.I("[polling-upgrader] Starting...")
	defer func() {
		u.Logger.I("[polling-upgrader] Done")
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apps, err := u.Apps.FindApps(ctx, server.AppsFilter{
		AutoUpgrade: lo.ToPtr(true),
	})
	if err != nil {
		u.Logger.E("[polling-upgrader] Failed to list apps: %+v", err)
		return
	}

	// Find apps with new images
	needsUpgraded := []internal.App{}
	for _, app := range apps {
		u.Logger.I("[polling-upgrader] Checking for update for %s (%s)...", app.Name, app.Image)
		currentDigest := app.ImageDigest
		newDigest, err := u.Runtime.PullLatest(ctx, app.Image)
		if err != nil {
			u.Logger.W("[polling-upgrader] Failed to pull latest image '%s' for '%s'", app.Image, app.Name)
			continue
		}
		if currentDigest == newDigest {
			u.Logger.V("[polling-upgrader] '%s' is already up to date", app.Name)
		} else {
			u.Logger.V("[polling-upgrader] '%s' has an update available", app.Name)
			app.ImageDigest = newDigest
			needsUpgraded = append(needsUpgraded, app)
		}
	}

	// Restart apps with new images
	plugins, err := u.Plugins.FindPlugins(ctx, server.PluginsFilter{})
	if err != nil {
		u.Logger.E("[polling-upgrader] Failed to list plugins: %+v", err)
		return
	}
	for _, app := range needsUpgraded {
		if err = u.upgradeApp(ctx, app, plugins); err != nil {
			u.Logger.W("[polling-upgrader] Failed to upgrade app: %+v", err)
		} else {
			u.Logger.I("[polling-upgrader] Upgraded '%s' to latest %s (%s)", app.Name, app.Image, app.ImageDigest)
		}
	}
}

func (u *PollingUpgrader) upgradeApp(ctx context.Context, app internal.App, plugins []internal.Plugin) error {
	route, err := u.Routes.FindRouteOrNil(ctx, server.RoutesFilter{
		AppID: &app.ID,
	})
	if err != nil {
		return err
	}

	env, err := u.Env.FindEnv(ctx, server.EnvFilter{AppID: &app.ID})
	if err != nil {
		return err
	}

	return u.Runtime.Restart(ctx, app, route, env, plugins)
}
