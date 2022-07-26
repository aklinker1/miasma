package graphql

import (
	"context"
	"errors"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
)

func (r *appResolver) getAppRoute(ctx context.Context, obj *internal.App) (*internal.Route, error) {
	if obj.Route != nil {
		return obj.Route, nil
	}
	route, err := inTx(ctx, r.DB.ReadonlyTx, zero.Route, func(tx server.Tx) (internal.Route, error) {
		return r.RouteRepo.GetOne(ctx, tx, server.RoutesFilter{
			AppID: &obj.ID,
		})
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &route, nil
}

func validateRouteInput(route *internal.RouteInput) error {
	if route == nil {
		return nil
	}
	if route.Host == nil && route.TraefikRule == nil {
		return errors.New("you must pass either a host or traefik rule")
	}
	return nil
}
