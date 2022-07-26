package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

func (r *queryResolver) Health(ctx context.Context) (*internal.Health, error) {
	return &internal.Health{
		Version: r.Version,
	}, nil
}

func (r *queryResolver) ListApps(ctx context.Context, page *int, size *int, showHidden *bool) ([]*internal.App, error) {
	apps, err := utils.InTx(ctx, r.DB.ReadonlyTx, nil, func(tx server.Tx) ([]internal.App, error) {
		return r.AppRepo.GetAll(ctx, tx, server.AppsFilter{
			IncludeHidden: showHidden,
			Pagination: &server.Pagination{
				Page: utils.ValueOr(page, 1),
				Size: utils.ValueOr(size, 10),
			},
		})
	})
	return utils.SafeReturn(lo.ToSlicePtr(apps), nil, err)
}

func (r *queryResolver) GetApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := utils.InTx(ctx, r.DB.ReadonlyTx, zero.App, func(tx server.Tx) (internal.App, error) {
		return r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
	})
	return utils.SafeReturn(&app, nil, err)
}

func (r *queryResolver) ListPlugins(ctx context.Context) ([]*internal.Plugin, error) {
	plugins, err := utils.InTx(ctx, r.DB.ReadonlyTx, nil, func(tx server.Tx) ([]internal.Plugin, error) {
		return r.PluginRepo.GetAll(ctx, tx, server.PluginsFilter{})
	})
	return utils.SafeReturn(lo.ToSlicePtr(plugins), nil, err)
}

func (r *queryResolver) GetPlugin(ctx context.Context, name internal.PluginName) (*internal.Plugin, error) {
	plugin, err := utils.InTx(ctx, r.DB.ReadonlyTx, zero.Plugin, func(tx server.Tx) (internal.Plugin, error) {
		return r.PluginRepo.GetOne(ctx, tx, server.PluginsFilter{
			Name: &name,
		})
	})
	return utils.SafeReturn(&plugin, nil, err)
}

func (r *queryResolver) Nodes(ctx context.Context) ([]*internal.Node, error) {
	nodes, err := r.RuntimeNodeRepo.GetAll(ctx, server.RuntimeNodesFilter{})
	return utils.SafeReturn(lo.ToSlicePtr(nodes), nil, err)
}

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
