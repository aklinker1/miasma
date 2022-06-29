package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

func (r *queryResolver) Health(ctx context.Context) (*internal.Health, error) {
	return &internal.Health{
		Version: r.Version,
	}, nil
}

func (r *queryResolver) ListApps(ctx context.Context, page *int32, size *int32, showHidden *bool) ([]*internal.App, error) {
	filter := server.AppsFilter{
		IncludeHidden: showHidden,
		Pagination: &server.Pagination{
			Page: utils.ValueOr(page, 1),
			Size: utils.ValueOr(size, 10),
		},
	}
	apps, err := r.Apps.FindApps(ctx, filter)
	return safeReturn(lo.ToSlicePtr(apps), nil, err)
}

func (r *queryResolver) GetApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.FindApp(ctx, server.AppsFilter{
		ID: &id,
	})
	return safeReturn(&app, nil, err)
}

func (r *queryResolver) ListPlugins(ctx context.Context) ([]*internal.Plugin, error) {
	plugins, err := r.Plugins.FindPlugins(ctx, server.PluginsFilter{})
	return safeReturn(lo.ToSlicePtr(plugins), nil, err)
}

func (r *queryResolver) GetPlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context) ([]*internal.Node, error) {
	nodes, err := r.Runtime.ListNodes(ctx)
	return safeReturn(lo.ToSlicePtr(nodes), nil, err)
}

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
