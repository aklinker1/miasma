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

func (r *healthResolver) Swarm(ctx context.Context, obj *internal.Health) (*internal.SwarmInfo, error) {
	swarm, err := r.Runtime.SwarmInfo(ctx)
	return safeReturn(swarm, nil, err)
}

// App returns gqlgen.AppResolver implementation.
func (r *Resolver) App() gqlgen.AppResolver { return &appResolver{r} }

// Health returns gqlgen.HealthResolver implementation.
func (r *Resolver) Health() gqlgen.HealthResolver { return &healthResolver{r} }

type appResolver struct{ *Resolver }
type healthResolver struct{ *Resolver }
