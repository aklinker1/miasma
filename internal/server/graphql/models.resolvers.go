package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

func (r *appResolver) Route(ctx context.Context, obj *internal.App) (*internal.Route, error) {
	route, err := r.getAppRoute(ctx, obj)
	return route, err
}

func (r *appResolver) SimpleRoute(ctx context.Context, obj *internal.App) (*string, error) {
	route, err := r.getAppRoute(ctx, obj)
	if err != nil {
		return nil, err
	}
	if route == nil {
		return nil, nil
	}

	traefik, err := inTx(ctx, r.DB.ReadonlyTx, zero.Plugin, func(tx server.Tx) (internal.Plugin, error) {
		return r.PluginRepo.GetTraefik(ctx, tx)
	})
	if err != nil {
		return nil, err
	}

	scheme := lo.Ternary(traefik.ConfigForTraefik().EnableHttps, "https", "http")
	if route.Host != nil && route.Path != nil {
		return lo.ToPtr(fmt.Sprintf("%s://%s/%s", scheme, *route.Host, *route.Path)), nil
	} else if route.Host != nil {
		return lo.ToPtr(fmt.Sprintf("%s://%s", scheme, *route.Host)), nil
	}
	return nil, nil
}

func (r *appResolver) AvailableAt(ctx context.Context, obj *internal.App, clusterIPAddress string) ([]string, error) {
	routes := []string{}
	simpleRoute, err := r.SimpleRoute(ctx, obj)
	if err != nil {
		return nil, err
	}
	if simpleRoute != nil {
		routes = append(routes, *simpleRoute)
	}

	runningService, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &obj.ID,
	})
	for _, port := range runningService.Endpoint.Ports {
		routes = append(routes, fmt.Sprintf("http://%s:%d", clusterIPAddress, port.PublishedPort))
	}
	if err != nil {
		r.Logger.W("Failed to load runtime information: %v", err)
	}

	return routes, nil
}

func (r *appResolver) Env(ctx context.Context, obj *internal.App) (map[string]interface{}, error) {
	env, err := inTx(ctx, r.DB.ReadonlyTx, nil, func(tx server.Tx) (internal.EnvMap, error) {
		return r.EnvRepo.Get(ctx, tx, server.EnvFilter{AppID: &obj.ID})
	})
	return safeReturn(utils.ToAnyMap(env), nil, err)
}

func (r *appResolver) Status(ctx context.Context, obj *internal.App) (internal.RuntimeStatus, error) {
	return r.getAppStatus(ctx, *obj)
}

func (r *appResolver) Instances(ctx context.Context, obj *internal.App) (*internal.AppInstances, error) {
	service, err := r.RuntimeServiceRepo.GetOne(ctx, server.RuntimeServicesFilter{
		AppID: &obj.ID,
	})
	if server.ErrorCode(err) == server.ENOTFOUND {
		return &internal.AppInstances{}, nil
	} else if err != nil {
		return nil, err
	}
	return &internal.AppInstances{
		Running: int(service.ServiceStatus.RunningTasks),
		Total:   int(service.ServiceStatus.DesiredTasks),
	}, nil
}

func (r *healthResolver) DockerVersion(ctx context.Context, obj *internal.Health) (string, error) {
	info, err := r.RuntimeRepo.Info(ctx)
	return info.ServerVersion, err
}

func (r *healthResolver) Cluster(ctx context.Context, obj *internal.Health) (*internal.ClusterInfo, error) {
	cluster, err := r.RuntimeRepo.ClusterInfo(ctx)
	if cluster == nil {
		return nil, err
	}
	return &internal.ClusterInfo{
		ID:          cluster.ID,
		JoinCommand: cluster.JoinTokens.Worker,
		CreatedAt:   cluster.CreatedAt,
		UpdatedAt:   cluster.UpdatedAt,
	}, nil
}

func (r *nodeResolver) Services(ctx context.Context, obj *internal.Node, showHidden *bool) ([]*internal.App, error) {
	services, err := r.RuntimeServiceRepo.GetAll(ctx, server.RuntimeServicesFilter{
		NodeID: &obj.ID,
	})
	if err != nil {
		return nil, err
	}

	return inTx(ctx, r.DB.ReadonlyTx, nil, func(tx server.Tx) ([]*internal.App, error) {
		apps := []*internal.App{}
		for _, service := range services {
			fmt.Println(service)
			app, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
				ID:            &service.AppID,
				IncludeHidden: showHidden,
			})
			if server.ErrorCode(err) == server.ENOTFOUND {
				// noop
			} else if err != nil {
				return nil, err
			} else {
				apps = append(apps, &app)
			}
		}
		return apps, nil
	})
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
