package cron

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

type PollingUpgrader struct {
	DB               server.DB
	Logger           server.Logger
	AppRepo          server.AppRepo
	RuntimeImageRepo server.RuntimeImageRepo
	PluginRepo       server.PluginRepo
	RuntimeService   *services.RuntimeService
}

func (u *PollingUpgrader) Cron() {
	u.Logger.I("[polling-upgrader] Starting...")
	defer func() {
		u.Logger.I("[polling-upgrader] Done")
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apps, err := utils.InTx(ctx, u.DB.ReadonlyTx, nil, func(tx server.Tx) ([]internal.App, error) {
		return u.AppRepo.GetAll(ctx, tx, server.AppsFilter{
			AutoUpgrade: lo.ToPtr(true),
		})
	})
	if err != nil {
		u.Logger.E("[polling-upgrader] Failed to list apps: %+v", err)
		return
	}

	// Find apps with new images
	needsUpgraded := []internal.App{}
	for _, app := range apps {
		u.Logger.I("[polling-upgrader] Checking for update for %s (%s)...", app.Name, app.Image)
		oldDigest := app.ImageDigest
		newDigest, err := u.RuntimeImageRepo.GetLatestDigest(ctx, app.Image)
		if err != nil {
			u.Logger.W("[polling-upgrader] Failed to pull latest image '%s' for '%s'", app.Image, app.Name)
			continue
		}
		if oldDigest == newDigest {
			u.Logger.V("[polling-upgrader] '%s' is already up to date", app.Name)
		} else {
			u.Logger.V("[polling-upgrader] '%s' has an update available", app.Name)
			app.ImageDigest = newDigest
			needsUpgraded = append(needsUpgraded, app)
		}
	}

	// Restart apps with new images
	plugins, err := utils.InTx(ctx, u.DB.ReadonlyTx, nil, func(tx server.Tx) ([]internal.Plugin, error) {
		return u.PluginRepo.GetAll(ctx, tx, server.PluginsFilter{})
	})
	if err != nil {
		u.Logger.E("[polling-upgrader] Failed to list plugins: %+v", err)
		return
	}
	for _, app := range needsUpgraded {
		if err := u.upgradeApp(ctx, app, plugins); err != nil {
			u.Logger.E("[polling-upgrader] Failed to upgrade app: %+v", err)
		} else {
			u.Logger.I("[polling-upgrader] Upgraded '%s' to latest %s (%s)", app.Name, app.Image, app.ImageDigest)
		}
	}
}

func (u *PollingUpgrader) upgradeApp(ctx context.Context, app internal.App, plugins []internal.Plugin) error {
	_, err := utils.InTx(ctx, u.DB.ReadonlyTx, zero.App, func(tx server.Tx) (internal.App, error) {
		updated, err := u.AppRepo.Update(ctx, tx, app)
		if err != nil {
			return zero.App, err
		}
		return updated, u.RuntimeService.RestartAppIfRunning(ctx, tx, services.PartialRuntimeServiceSpec{
			App:        updated,
			HasPlugins: true,
			Plugins:    plugins,
		})
	})
	return err
}
