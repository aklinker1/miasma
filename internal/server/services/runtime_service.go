package services

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
)

type RuntimeService struct {
	DB                 server.DB
	Logger             server.Logger
	AppRepo            server.AppRepo
	RouteRepo          server.RouteRepo
	EnvRepo            server.EnvRepo
	PluginRepo         server.PluginRepo
	RuntimeServiceRepo server.RuntimeServiceRepo
	AppService         *AppService
}

// Struct containing all the required data to "start" and app. Only include the data you already know
type PartialRuntimeServiceSpec struct {
	App        internal.App
	HasPlugins bool
	Plugins    []internal.Plugin
	HasEnv     bool
	Env        internal.EnvMap
	HasRoute   bool
	Route      *internal.Route
}

func (r *RuntimeService) StartApp(ctx context.Context, tx server.Tx, partialInput PartialRuntimeServiceSpec) error {
	spec, err := r.getRuntimeServiceSpec(ctx, tx, partialInput)
	if err != nil {
		return err
	}
	err = r.RuntimeServiceRepo.Create(ctx, spec)
	return err
}

func (r *RuntimeService) StopApp(ctx context.Context, tx server.Tx, partialInput PartialRuntimeServiceSpec) error {
	existing, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &partialInput.App.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		r.Logger.W("Stopping service that is already stopped")
		return nil
	}
	if err != nil {
		return err
	}
	_, err = r.RuntimeServiceRepo.Remove(ctx, existing)
	return err
}

func (r *RuntimeService) RestartApp(ctx context.Context, tx server.Tx, partialInput PartialRuntimeServiceSpec) error {
	// Stop
	existing, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &partialInput.App.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		r.Logger.W("Restarting service that is already stopped")
	} else if err != nil {
		return err
	} else {
		_, err = r.RuntimeServiceRepo.Remove(ctx, existing)
		if err != nil {
			return err
		}
	}

	// Start
	spec, err := r.getRuntimeServiceSpec(ctx, tx, partialInput)
	if err != nil {
		return err
	}
	err = r.RuntimeServiceRepo.Create(ctx, spec)
	return err
}

func (r *RuntimeService) RestartAppIfRunning(ctx context.Context, tx server.Tx, partialInput PartialRuntimeServiceSpec) error {
	status, err := r.AppService.GetAppStatus(ctx, partialInput.App)
	if err != nil {
		return err
	}
	if status == internal.RuntimeStatusRunning {
		return r.RestartApp(ctx, tx, partialInput)
	}
	return nil
}

func (r *RuntimeService) getRuntimeServiceSpec(
	ctx context.Context,
	tx server.Tx,
	partial PartialRuntimeServiceSpec,
) (server.RuntimeServiceSpec, error) {
	data := server.RuntimeServiceSpec{
		App: partial.App,
	}

	// Env
	if partial.HasEnv {
		data.Env = partial.Env
	} else {
		env, err := r.EnvRepo.Get(ctx, tx, server.EnvFilter{AppID: data.App.ID})
		if err != nil {
			return zero.RuntimeServiceSpec, err
		}
		data.Env = env
	}

	// Route
	if partial.HasRoute {
		data.Route = partial.Route
	} else {
		route, err := r.RouteRepo.GetOne(ctx, tx, server.RoutesFilter{
			AppID: &data.App.ID,
		})
		if server.ErrorCode(err) == server.ENOTFOUND {
			data.Route = nil
		} else if err != nil {
			return zero.RuntimeServiceSpec, err
		} else {
			data.Route = &route
		}
	}

	// Plugins
	if partial.HasPlugins {
		data.Plugins = partial.Plugins
	} else {
		plugins, err := r.PluginRepo.GetAll(ctx, tx, server.PluginsFilter{})
		if err != nil {
			return zero.RuntimeServiceSpec, err
		}
		data.Plugins = plugins
	}

	return data, nil
}

func (r *RuntimeService) RestartAllAppsIfRunning(ctx context.Context, tx server.Tx) error {
	apps, err := r.AppRepo.GetAll(ctx, tx, server.AppsFilter{})
	if err != nil {
		return err
	}

	plugins, err := r.PluginRepo.GetAll(ctx, tx, server.PluginsFilter{})
	if err != nil {
		return err
	}

	for _, app := range apps {
		err = r.RestartAppIfRunning(ctx, tx, PartialRuntimeServiceSpec{
			App:        app,
			HasPlugins: true,
			Plugins:    plugins,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
