package services

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
)

type AppService struct {
	DB                 server.DB
	Logger             server.Logger
	AppRepo            server.AppRepo
	RouteRepo          server.RouteRepo
	EnvRepo            server.EnvRepo
	PluginRepo         server.PluginRepo
	RuntimeServiceRepo server.RuntimeServiceRepo
}

func (s *AppService) GetAppRoute(ctx context.Context, obj *internal.App) (*internal.Route, error) {
	if obj.Route != nil {
		return obj.Route, nil
	}
	route, err := utils.InTx(ctx, s.DB.ReadonlyTx, zero.Route, func(tx server.Tx) (internal.Route, error) {
		return s.RouteRepo.GetOne(ctx, tx, server.RoutesFilter{
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

func (s *AppService) GetAppStatus(ctx context.Context, app internal.App) (internal.RuntimeStatus, error) {
	_, err := s.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &app.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		return internal.RuntimeStatusStopped, nil
	} else if err != nil {
		return internal.RuntimeStatusStopped, err
	}
	return internal.RuntimeStatusRunning, nil
}

func (s *AppService) DeleteAppCascade(ctx context.Context, tx server.Tx, app internal.App) (internal.App, error) {
	deleted, err := s.AppRepo.Delete(ctx, tx, app)
	if err != nil {
		return zero.App, err
	}

	// Routes
	routes, err := s.RouteRepo.GetAll(ctx, tx, server.RoutesFilter{
		AppID: &deleted.ID,
	})
	if err != nil {
		return zero.App, err
	}
	for _, route := range routes {
		_, err = s.RouteRepo.Delete(ctx, tx, route)
		if err != nil {
			return zero.App, err
		}
	}

	// Clear Env
	_, err = s.EnvRepo.Set(ctx, tx, deleted.ID, internal.EnvMap{})
	if err != nil {
		return zero.App, err
	}

	return deleted, nil
}
