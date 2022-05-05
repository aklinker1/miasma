package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
)

func (r *healthResolver) DockerVersion(ctx context.Context, obj *internal.Health) (string, error) {
	return r.Runtime.Version(ctx)
}

func (r *healthResolver) Swarm(ctx context.Context, obj *internal.Health) (*internal.SwarmInfo, error) {
	return safeReturn(r.Runtime.SwarmInfo(ctx))
}

// Health returns gqlgen.HealthResolver implementation.
func (r *Resolver) Health() gqlgen.HealthResolver { return &healthResolver{r} }

type healthResolver struct{ *Resolver }
