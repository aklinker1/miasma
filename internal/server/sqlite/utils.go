package sqlite

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

type startParamKnowns struct {
	app        internal.App
	knownApp   bool
	route      *internal.Route
	knownRoute bool
	env        internal.EnvMap
	knownEnv   bool
}

func findStartParams(ctx context.Context, tx server.Tx, appID string, known startParamKnowns) (internal.App, *internal.Route, internal.EnvMap, error) {
	var err error
	app := known.app
	if !known.knownApp {
		app, err = findApp(ctx, tx, server.AppsFilter{ID: &appID})
	}
	if err != nil {
		return known.app, known.route, known.env, err
	}

	route := known.route
	if !known.knownApp {
		route, err = findRouteOrNil(ctx, tx, server.RoutesFilter{AppID: &appID})
	}
	if err != nil {
		return known.app, known.route, known.env, err
	}

	env := known.env
	if !known.knownApp {
		env, err = findEnvMap(ctx, tx, server.EnvFilter{AppID: &appID})
	}
	if err != nil {
		return known.app, known.route, known.env, err
	}

	return app, route, env, nil
}
