package graphql

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/samber/lo"
)

func safeReturn[T any](value T, fallback T, err error) (T, error) {
	if err != nil {
		return fallback, err
	} else {
		return value, err
	}
}

func (r *appResolver) getAppRoute(ctx context.Context, obj *internal.App) (*internal.AppRouting, error) {
	traefik, err := r.Plugins.FindPlugin(ctx, server.PluginsFilter{
		Name: lo.ToPtr(internal.PluginNameTraefik),
	})
	if err != nil {
		return nil, err
	}
	if !traefik.Enabled {
		return nil, nil
	}
	if obj.Routing != nil {
		return obj.Routing, nil
	}

	route, err := r.Routes.FindRoute(ctx, server.RoutesFilter{
		AppID: &obj.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &route, nil
	}
}
