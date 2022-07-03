package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

type startParamKnowns struct {
	app          internal.App
	knownApp     bool
	route        *internal.Route
	knownRoute   bool
	env          internal.EnvMap
	knownEnv     bool
	plugins      []internal.Plugin
	knownPlugins bool
}

func findStartParams(ctx context.Context, tx server.Tx, appID string, known startParamKnowns) (internal.App, *internal.Route, internal.EnvMap, []internal.Plugin, error) {
	var err error
	app := known.app
	if !known.knownApp {
		app, err = findApp(ctx, tx, server.AppsFilter{ID: &appID})
	}
	if err != nil {
		return known.app, known.route, known.env, known.plugins, err
	}

	route := known.route
	if !known.knownApp {
		route, err = findRouteOrNil(ctx, tx, server.RoutesFilter{AppID: &appID})
	}
	if err != nil {
		return known.app, known.route, known.env, known.plugins, err
	}

	env := known.env
	if !known.knownApp {
		env, err = findEnvMap(ctx, tx, server.EnvFilter{AppID: &appID})
	}
	if err != nil {
		return known.app, known.route, known.env, known.plugins, err
	}

	plugins := known.plugins
	if !known.knownPlugins {
		plugins, err = findPlugins(ctx, tx, server.PluginsFilter{})
	}
	if err != nil {
		return known.app, known.route, known.env, known.plugins, err
	}

	return app, route, env, plugins, nil
}
