package graphql

import (
	"context"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
)

func isSwarmNotInitializedError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "This node is not a swarm manager")
}

// Struct containing all the required data to "start" and app. Only include the data you already know
type partialRuntimeServiceSpec struct {
	app        internal.App
	hasPlugins bool
	plugins    []internal.Plugin
	hasEnv     bool
	env        internal.EnvMap
	hasRoute   bool
	route      *internal.Route
}

func (r *Resolver) startApp(ctx context.Context, tx server.Tx, partialInput partialRuntimeServiceSpec) error {
	spec, err := r.getRuntimeServiceSpec(ctx, tx, partialInput)
	if err != nil {
		return err
	}
	_, err = r.RuntimeServiceRepo.Create(ctx, spec)
	return err
}

func (r *Resolver) stopApp(ctx context.Context, tx server.Tx, partialInput partialRuntimeServiceSpec) error {
	existing, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &partialInput.app.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		r.Logger.W("Stopping service that is already stopped")
		return nil
	}
	if err != nil {
		return err
	}
	_, err = r.RuntimeServiceRepo.Delete(ctx, existing)
	return err
}

func (r *Resolver) restartApp(ctx context.Context, tx server.Tx, partialInput partialRuntimeServiceSpec) error {
	// Stop
	existing, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &partialInput.app.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		r.Logger.W("Restarting service that is already stopped")
	} else if err != nil {
		return err
	} else {
		_, err = r.RuntimeServiceRepo.Delete(ctx, existing)
		if err != nil {
			return err
		}
	}

	// Start
	spec, err := r.getRuntimeServiceSpec(ctx, tx, partialInput)
	if err != nil {
		return err
	}
	_, err = r.RuntimeServiceRepo.Create(ctx, spec)
	return err
}

func (r *Resolver) restartAppIfRunning(ctx context.Context, tx server.Tx, partialInput partialRuntimeServiceSpec) error {
	status, err := r.getAppStatus(ctx, partialInput.app)
	if err != nil {
		return err
	}
	if status == internal.RuntimeStatusRunning {
		return r.restartApp(ctx, tx, partialInput)
	}
	return nil
}

func (r *Resolver) getRuntimeServiceSpec(
	ctx context.Context,
	tx server.Tx,
	partial partialRuntimeServiceSpec,
) (server.RuntimeServiceSpec, error) {
	data := server.RuntimeServiceSpec{
		App: partial.app,
	}

	// Env
	if partial.hasEnv {
		data.Env = partial.env
	} else {
		env, err := r.EnvRepo.Get(ctx, tx, server.EnvFilter{
			AppID: &data.App.ID,
		})
		if err != nil {
			return zero.RuntimeServiceSpec, err
		}
		data.Env = env
	}

	// Route
	if partial.hasRoute {
		data.Route = partial.route
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
	if partial.hasPlugins {
		data.Plugins = partial.plugins
	} else {
		plugins, err := r.PluginRepo.GetAll(ctx, tx, server.PluginsFilter{})
		if err != nil {
			return zero.RuntimeServiceSpec, err
		}
		data.Plugins = plugins
	}

	return data, nil
}

func (r *Resolver) restartAllAppsIfRunning(ctx context.Context, tx server.Tx) error {
	apps, err := r.AppRepo.GetAll(ctx, tx, server.AppsFilter{})
	if err != nil {
		return err
	}

	plugins, err := r.PluginRepo.GetAll(ctx, tx, server.PluginsFilter{})
	if err != nil {
		return err
	}

	for _, app := range apps {
		err = r.restartAppIfRunning(ctx, tx, partialRuntimeServiceSpec{
			app:        app,
			hasPlugins: true,
			plugins:    plugins,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
