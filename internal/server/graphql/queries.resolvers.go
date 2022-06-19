package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/utils"
)

func (r *queryResolver) Health(ctx context.Context) (*internal.Health, error) {
	return &internal.Health{
		Version: r.Version,
	}, nil
}

func (r *queryResolver) ListApps(ctx context.Context, page *int32, size *int32, showHidden *bool) ([]*internal.App, error) {
	filter := server.GetAppFilters{
		IncludeHidden: utils.BoolOr(showHidden, false),
		Pagination: server.Pagination{
			Page: utils.Int32Or(page, 1),
			Size: utils.Int32Or(size, 10),
		},
	}
	apps, err := r.Apps.FindApps(ctx, filter)
	return safeReturn(apps, nil, err)
}

func (r *queryResolver) GetApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.FindAppByName(ctx, appName)
	return safeReturn(&app, nil, err)
}

func (r *queryResolver) ListPlugins(ctx context.Context) ([]*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAppRouting(ctx context.Context, appID string) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
