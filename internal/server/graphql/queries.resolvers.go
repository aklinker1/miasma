package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
)

func (r *queryResolver) Health(ctx context.Context) (*internal.Health, error) {
	swarm, err := r.Runtime.SwarmInfo(ctx)
	if err != nil {
		return nil, err
	}
	dockerVersion, err := r.Runtime.Version(ctx)
	if err != nil {
		return nil, err
	}
	return &internal.Health{
		Version:       r.Version,
		DockerVersion: dockerVersion,
		Swarm:         swarm,
	}, nil
}

func (r *queryResolver) ListApps(ctx context.Context, page *int32, size *int32, showHidden *bool) ([]internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetApp(ctx context.Context, appName string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListPlugins(ctx context.Context) ([]internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAppRouting(ctx context.Context, appName string) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
