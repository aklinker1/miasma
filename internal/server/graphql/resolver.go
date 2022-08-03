package graphql

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Version string
	DB      server.DB
	Logger  server.Logger

	AppRepo            server.AppRepo
	PluginRepo         server.PluginRepo
	EnvRepo            server.EnvRepo
	RouteRepo          server.RouteRepo
	RuntimeRepo        server.RuntimeRepo
	RuntimeServiceRepo server.RuntimeServiceRepo
	RuntimeNodeRepo    server.RuntimeNodeRepo
	RuntimeTaskRepo    server.RuntimeTaskRepo
	RuntimeImageRepo   server.RuntimeImageRepo

	AppService     *services.AppService
	PluginService  *services.PluginService
	RuntimeService *services.RuntimeService
}

func (r *Resolver) getApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := utils.InTx(ctx, r.DB.ReadonlyTx, zero.App, func(tx server.Tx) (internal.App, error) {
		return r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
	})
	return utils.SafeReturn(&app, nil, err)
}

func (r *Resolver) getNode(ctx context.Context, id string) (*internal.Node, error) {
	node, err := r.RuntimeNodeRepo.GetOne(ctx, server.RuntimeNodesFilter{
		ID: &id,
	})
	return utils.SafeReturn(&node, nil, err)
}
