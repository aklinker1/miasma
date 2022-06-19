package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/samber/lo"
)

func (r *appResolver) Routing(ctx context.Context, obj *internal.App) (*internal.AppRouting, error) {
	route, err := r.Routes.FindRoute(ctx, server.RoutesFilter{
		AppID: &obj.ID,
	})
	return safeReturn(&route, nil, err)
}

func (r *appResolver) SimpleRoute(ctx context.Context, obj *internal.App) (*string, error) {
	route, err := r.Routes.FindRoute(ctx, server.RoutesFilter{
		AppID: &obj.ID,
	})
	if err != nil {
		return nil, err
	}
	if route.TraefikRule != nil {
		return route.TraefikRule, nil
	} else {
		return lo.ToPtr(fmt.Sprintf("%s/%s", *route.Host, *route.Path)), nil
	}
}

func (r *appResolver) Status(ctx context.Context, obj *internal.App) (string, error) {
	return "TODO", nil
}

func (r *appResolver) Instances(ctx context.Context, obj *internal.App) (string, error) {
	return "TODO", nil
}

func (r *healthResolver) DockerVersion(ctx context.Context, obj *internal.Health) (string, error) {
	return r.Runtime.Version(ctx)
}

func (r *healthResolver) Cluster(ctx context.Context, obj *internal.Health) (*internal.ClusterInfo, error) {
	swarm, err := r.Runtime.ClusterInfo(ctx)
	return safeReturn(swarm, nil, err)
}

// App returns gqlgen.AppResolver implementation.
func (r *Resolver) App() gqlgen.AppResolver { return &appResolver{r} }

// Health returns gqlgen.HealthResolver implementation.
func (r *Resolver) Health() gqlgen.HealthResolver { return &healthResolver{r} }

type appResolver struct{ *Resolver }
type healthResolver struct{ *Resolver }
