package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
)

func (r *appResolver) Status(ctx context.Context, obj *internal.App) (string, error) {
	return "TODO", nil
}

func (r *appResolver) Instances(ctx context.Context, obj *internal.App) (string, error) {
	return "TODO", nil
}

func (r *healthResolver) DockerVersion(ctx context.Context, obj *internal.Health) (string, error) {
	return r.Runtime.Version(ctx)
}

// App returns gqlgen.AppResolver implementation.
func (r *Resolver) App() gqlgen.AppResolver { return &appResolver{r} }

// Health returns gqlgen.HealthResolver implementation.
func (r *Resolver) Health() gqlgen.HealthResolver { return &healthResolver{r} }

type appResolver struct{ *Resolver }
type healthResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *healthResolver) ClusterInfo(ctx context.Context, obj *internal.Health) (*internal.ClusterInfo, error) {
	swarm, err := r.Runtime.ClusterInfo(ctx)
	return safeReturn(swarm, nil, err)
}
