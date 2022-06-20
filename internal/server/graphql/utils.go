package graphql

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
)

func safeReturn[T any](value T, fallback T, err error) (T, error) {
	if err != nil {
		return fallback, err
	} else {
		return value, err
	}
}

func getAppRoute(ctx context.Context, routes server.RouteService, obj *internal.App) (*internal.AppRouting, error) {
	if obj.Routing != nil {
		return obj.Routing, nil
	}

	route, err := routes.FindRoute(ctx, server.RoutesFilter{
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
