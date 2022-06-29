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

func (r *appResolver) Route(ctx context.Context, obj *internal.App) (*internal.Route, error) {
	return r.getAppRoute(ctx, obj)
}

func (r *appResolver) SimpleRoute(ctx context.Context, obj *internal.App) (*string, error) {
	route, err := r.getAppRoute(ctx, obj)
	if err != nil {
		return nil, err
	}
	if route == nil {
		return nil, nil
	}

	if route.TraefikRule != nil {
		return route.TraefikRule, nil
	} else if route.Host != nil && route.Path != nil {
		return lo.ToPtr(fmt.Sprintf("%s/%s", *route.Host, *route.Path)), nil
	} else if route.Host != nil {
		return route.Host, nil
	}
	return nil, nil
}

func (r *appResolver) Env(ctx context.Context, obj *internal.App) (map[string]interface{}, error) {
	env, err := r.EnvService.FindEnv(ctx, server.EnvFilter{AppID: &obj.ID})
	return safeReturn(utils.ToAnyMap(env), nil, err)
}

func (r *appResolver) Status(ctx context.Context, obj *internal.App) (string, error) {
	info, err := r.Runtime.GetRuntimeAppInfo(ctx, *obj)
	if server.ErrorCode(err) == server.ENOTFOUND {
		return "stopped", nil
	} else if err != nil {
		return "", err
	}
	return info.Status, nil
}

func (r *appResolver) Instances(ctx context.Context, obj *internal.App) (*internal.AppInstances, error) {
	info, err := r.Runtime.GetRuntimeAppInfo(ctx, *obj)
	if server.ErrorCode(err) == server.ENOTFOUND {
		return &internal.AppInstances{}, nil
	} else if err != nil {
		return nil, err
	}
	return &info.Instances, nil
}

func (r *healthResolver) DockerVersion(ctx context.Context, obj *internal.Health) (string, error) {
	return r.Runtime.Version(ctx)
}

func (r *healthResolver) Cluster(ctx context.Context, obj *internal.Health) (*internal.ClusterInfo, error) {
	swarm, err := r.Runtime.ClusterInfo(ctx)
	return safeReturn(swarm, nil, err)
}

func (r *nodeResolver) Services(ctx context.Context, obj *internal.Node) ([]*internal.RunningContainer, error) {
	services, err := r.Runtime.ListServices(ctx, server.ListServicesFilter{
		NodeID: &obj.ID,
	})
	return safeReturn(lo.ToSlicePtr(services), nil, err)
}

// App returns gqlgen.AppResolver implementation.
func (r *Resolver) App() gqlgen.AppResolver { return &appResolver{r} }

// Health returns gqlgen.HealthResolver implementation.
func (r *Resolver) Health() gqlgen.HealthResolver { return &healthResolver{r} }

// Node returns gqlgen.NodeResolver implementation.
func (r *Resolver) Node() gqlgen.NodeResolver { return &nodeResolver{r} }

type appResolver struct{ *Resolver }
type healthResolver struct{ *Resolver }
type nodeResolver struct{ *Resolver }
