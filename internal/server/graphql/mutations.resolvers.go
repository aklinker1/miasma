package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateApp(ctx context.Context, input internal.AppInput) (*internal.App, error) {
	if input.Name = strings.TrimSpace(input.Name); input.Name == "" {
		return nil, &server.Error{
			Code:    server.EINVALID,
			Message: "App name cannot be empty",
			Op:      "createApp",
		}
	}
	if input.Image = strings.TrimSpace(input.Image); input.Image == "" {
		return nil, &server.Error{
			Code:    server.EINVALID,
			Message: "App image cannot be empty",
			Op:      "createApp",
		}
	}

	a := internal.App{
		CreatedAt:      time.Now(),
		Name:           input.Name,
		Group:          input.Group,
		Image:          input.Image,
		Hidden:         utils.ValueOr(input.Hidden, false),
		TargetPorts:    input.TargetPorts,
		PublishedPorts: input.PublishedPorts,
		Placement:      input.Placement,
		Volumes: lo.Map(input.Volumes, func(v *internal.BoundVolumeInput, _ int) *internal.BoundVolume {
			return &internal.BoundVolume{
				Target: v.Target,
				Source: v.Source,
			}
		}),
		Networks: input.Networks,
		Command:  input.Command,
	}

	created, err := r.Apps.Create(ctx, a)
	return safeReturn(&created, nil, err)
}

func (r *mutationResolver) EditApp(ctx context.Context, id string, changes map[string]interface{}) (*internal.App, error) {
	newApp, err := r.Apps.FindApp(ctx, server.AppsFilter{
		ID: &id,
	})
	if err != nil {
		return nil, err
	}
	gqlgen.ApplyChanges(changes, &newApp)

	// Grab new image from changes
	var newImage *string
	newImageStr, ok := changes["image"].(string)
	if ok {
		newImage = &newImageStr
	}

	updated, err := r.Apps.Update(ctx, newApp, newImage)
	return safeReturn(&updated, nil, err)
}

func (r *mutationResolver) DeleteApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.Delete(ctx, id)
	return safeReturn(&app, nil, err)
}

func (r *mutationResolver) StartApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.FindApp(ctx, server.AppsFilter{ID: &id})
	if err != nil {
		return nil, err
	}

	route, err := r.Routes.FindRouteOrNil(ctx, server.RoutesFilter{
		AppID: &id,
	})
	if err != nil {
		return nil, err
	}

	env, err := r.EnvService.FindEnv(ctx, server.EnvFilter{
		AppID: &id,
	})
	if err != nil {
		return nil, err
	}

	err = r.Runtime.Start(ctx, app, route, env)
	return safeReturn(&app, nil, err)
}

func (r *mutationResolver) StopApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.FindApp(ctx, server.AppsFilter{ID: &id})
	if err != nil {
		return nil, err
	}

	err = r.Runtime.Stop(ctx, app)
	return safeReturn(&app, nil, err)
}

func (r *mutationResolver) RestartApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.FindApp(ctx, server.AppsFilter{ID: &id})
	if err != nil {
		return nil, err
	}
	route, err := r.Routes.FindRouteOrNil(ctx, server.RoutesFilter{AppID: &id})
	if err != nil {
		return nil, err
	}
	env, err := r.EnvService.FindEnv(ctx, server.EnvFilter{AppID: &id})
	if err != nil {
		return nil, err
	}

	err = r.Runtime.Restart(ctx, app, route, env)
	return safeReturn(&app, nil, err)
}

func (r *mutationResolver) UpgradeApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.FindApp(ctx, server.AppsFilter{ID: &id})
	if err != nil {
		return nil, err
	}
	updated, err := r.Apps.Update(ctx, app, &app.Image)
	return safeReturn(&updated, nil, err)
}

func (r *mutationResolver) EnablePlugin(ctx context.Context, name internal.PluginName, config map[string]interface{}) (*internal.Plugin, error) {
	plugin, err := r.Plugins.FindPlugin(ctx, server.PluginsFilter{
		Name: &name,
	})
	if err != nil {
		return nil, err
	}

	updated, err := r.Plugins.EnablePlugin(ctx, plugin, config)
	return safeReturn(&updated, nil, err)
}

func (r *mutationResolver) DisablePlugin(ctx context.Context, name internal.PluginName) (*internal.Plugin, error) {
	plugin, err := r.Plugins.FindPlugin(ctx, server.PluginsFilter{
		Name: &name,
	})
	if err != nil {
		return nil, err
	}

	updated, err := r.Plugins.DisablePlugin(ctx, plugin)
	return safeReturn(&updated, nil, err)
}

func (r *mutationResolver) SetAppEnv(ctx context.Context, appID string, newEnv map[string]interface{}) (map[string]interface{}, error) {
	created, err := r.EnvService.SetAppEnv(ctx, appID, utils.ToEnvMap(newEnv))
	return safeReturn(utils.ToAnyMap(created), nil, err)
}

func (r *mutationResolver) SetAppRoute(ctx context.Context, appID string, route *internal.RouteInput) (*internal.Route, error) {
	if route.Host == nil && route.TraefikRule == nil {
		return nil, errors.New("You must pass either a host or traefik rule")
	}

	existing, err := r.Routes.FindRoute(ctx, server.RoutesFilter{AppID: &appID})

	if server.ErrorCode(err) == server.ENOTFOUND {
		// Create a new route
		created, err := r.Routes.Create(ctx, internal.Route{
			AppID:       appID,
			Host:        route.Host,
			Path:        route.Path,
			TraefikRule: route.TraefikRule,
		})
		return safeReturn(&created, nil, err)
	} else if err == nil {
		// Update existing route
		existing.Host = route.Host
		existing.Path = route.Path
		existing.TraefikRule = route.TraefikRule
		updated, err := r.Routes.Update(ctx, existing)
		return safeReturn(&updated, nil, err)
	}

	return nil, err
}

func (r *mutationResolver) RemoveAppRoute(ctx context.Context, appID string) (*internal.Route, error) {
	existing, err := r.Routes.FindRoute(ctx, server.RoutesFilter{AppID: &appID})
	if err != nil {
		return nil, err
	}
	deleted, err := r.Routes.Delete(ctx, existing)
	return safeReturn(&deleted, nil, err)
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
