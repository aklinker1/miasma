package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/server/services"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateApp(ctx context.Context, input internal.AppInput) (*internal.App, error) {
	err := validateAppInput(input)
	if err != nil {
		return nil, err
	}

	// Define the app
	newApp := internal.App{
		CreatedAt:      time.Now(),
		Name:           input.Name,
		Group:          input.Group,
		Image:          input.Image,
		AutoUpgrade:    utils.ValueOr(input.AutoUpgrade, true),
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

	// Grab the image digest
	newDigest, err := r.RuntimeImageRepo.GetLatestDigest(ctx, newApp.Image)
	if err != nil {
		return nil, err
	}
	newApp.ImageDigest = newDigest

	created, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		// Save the app
		created, err := r.AppRepo.Create(ctx, tx, newApp)
		if err != nil {
			return zero.App, err
		}

		// Start the app
		err = r.RuntimeService.StartApp(ctx, tx, services.PartialRuntimeServiceSpec{
			App: created,
		})
		return created, err
	})
	return utils.SafeReturn(&created, nil, err)
}

func (r *mutationResolver) EditApp(ctx context.Context, id string, changes map[string]interface{}) (*internal.App, error) {
	updated, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		// Apply changes
		newApp, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
		if err != nil {
			return zero.App, err
		}
		err = gqlgen.ApplyChanges(changes, &newApp)
		if err != nil {
			return zero.App, err
		}

		// Pull new image digest if image is included
		if newImage, ok := changes["image"].(string); ok {
			newDigest, err := r.RuntimeImageRepo.GetLatestDigest(ctx, newImage)
			if err != nil {
				return zero.App, err
			}
			newApp.ImageDigest = newDigest
		}

		// Save changes
		updated, err := r.AppRepo.Update(ctx, tx, newApp)
		if err != nil {
			return zero.App, err
		}

		// Restart the app
		return updated, r.RuntimeService.RestartAppIfRunning(ctx, tx, services.PartialRuntimeServiceSpec{
			App: updated,
		})
	})
	return utils.SafeReturn(&updated, nil, err)
}

func (r *mutationResolver) DeleteApp(ctx context.Context, id string) (*internal.App, error) {
	deleted, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		app, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
		if err != nil {
			return zero.App, err
		}
		deleted, err := r.AppRepo.Delete(ctx, tx, app)
		if err != nil {
			return zero.App, err
		}
		return deleted, r.RuntimeService.StopApp(ctx, tx, services.PartialRuntimeServiceSpec{
			App: deleted,
		})
	})
	return utils.SafeReturn(&deleted, nil, err)
}

func (r *mutationResolver) StartApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		app, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
		if err != nil {
			return zero.App, err
		}
		return app, r.RuntimeService.StartApp(ctx, tx, services.PartialRuntimeServiceSpec{
			App: app,
		})
	})
	return utils.SafeReturn(&app, nil, err)
}

func (r *mutationResolver) StopApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		app, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
		if err != nil {
			return zero.App, err
		}
		return app, r.RuntimeService.StopApp(ctx, tx, services.PartialRuntimeServiceSpec{
			App: app,
		})
	})
	return utils.SafeReturn(&app, nil, err)
}

func (r *mutationResolver) RestartApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		app, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
		if err != nil {
			return zero.App, err
		}
		return app, r.RuntimeService.RestartApp(ctx, tx, services.PartialRuntimeServiceSpec{
			App: app,
		})
	})
	return utils.SafeReturn(&app, nil, err)
}

func (r *mutationResolver) UpgradeApp(ctx context.Context, id string) (*internal.App, error) {
	upgraded, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.App, func(tx server.Tx) (internal.App, error) {
		// Grab the old digest
		app, err := r.AppRepo.GetOne(ctx, tx, server.AppsFilter{
			ID: &id,
		})
		if err != nil {
			return zero.App, err
		}
		oldDigest := app.ImageDigest

		// Grab the latest digest
		newDigest, err := r.RuntimeImageRepo.GetLatestDigest(ctx, app.Image)
		if err != nil {
			return zero.App, err
		}

		// Compare and save
		if oldDigest == newDigest {
			return app, nil
		}
		app.ImageDigest = newDigest
		upgraded, err := r.AppRepo.Update(ctx, tx, app)
		if err != nil {
			return zero.App, err
		}

		// Restart the app
		return upgraded, r.RuntimeService.RestartAppIfRunning(ctx, tx, services.PartialRuntimeServiceSpec{
			App: upgraded,
		})
	})
	return utils.SafeReturn(&upgraded, nil, err)
}

func (r *mutationResolver) EnablePlugin(ctx context.Context, name internal.PluginName, config map[string]interface{}) (*internal.Plugin, error) {
	enabled, err := r.PluginService.TogglePlugin(ctx, true, name, config)
	return utils.SafeReturn(&enabled, nil, err)
}

func (r *mutationResolver) DisablePlugin(ctx context.Context, name internal.PluginName) (*internal.Plugin, error) {
	disabled, err := r.PluginService.TogglePlugin(ctx, false, name, map[string]any{})
	return utils.SafeReturn(&disabled, nil, err)
}

func (r *mutationResolver) SetAppEnv(ctx context.Context, appID string, newEnv map[string]interface{}) (map[string]interface{}, error) {
	created, err := utils.InTx(ctx, r.DB.ReadWriteTx, nil, func(tx server.Tx) (internal.EnvMap, error) {
		return r.EnvRepo.Set(ctx, tx, appID, utils.ToEnvMap(newEnv))
	})
	return utils.SafeReturn(utils.ToAnyMap(created), nil, err)
}

func (r *mutationResolver) SetAppRoute(ctx context.Context, appID string, route *internal.RouteInput) (*internal.Route, error) {
	err := validateRouteInput(route)
	if err != nil {
		return nil, err
	}

	updated, err := utils.InTx(ctx, r.DB.ReadWriteTx, zero.Route, func(tx server.Tx) (internal.Route, error) {
		existing, err := r.RouteRepo.GetOne(ctx, tx, server.RoutesFilter{
			AppID: &appID,
		})
		if server.ErrorCode(err) == server.ENOTFOUND {
			if route == nil {
				// No route found, and set to nil, so nothing to do
				return existing, nil
			} else {
				// No route found, but one was passed in, so we need to create it
				return r.RouteRepo.Create(ctx, tx, internal.Route{
					AppID:       appID,
					Host:        route.Host,
					Path:        route.Path,
					TraefikRule: route.TraefikRule,
				})
			}
		}
		if err != nil {
			return zero.Route, err
		}

		if route == nil {
			// Found a route, but passed in nil, so delete it
			return r.RouteRepo.Delete(ctx, tx, existing)
		}

		// Found a route, passed in something new, so update it
		existing.Host = route.Host
		existing.Path = route.Path
		existing.TraefikRule = route.TraefikRule
		return r.RouteRepo.Update(ctx, tx, existing)
	})

	return utils.SafeReturn(&updated, nil, err)
}

func (r *mutationResolver) RemoveAppRoute(ctx context.Context, appID string) (*internal.Route, error) {
	return utils.InTx(ctx, r.DB.ReadWriteTx, nil, func(tx server.Tx) (*internal.Route, error) {
		route, err := r.RouteRepo.GetOne(ctx, tx, server.RoutesFilter{
			AppID: &appID,
		})
		if server.ErrorCode(err) == server.ENOTFOUND {
			return nil, nil
		} else if err != nil {
			return &zero.Route, err
		}
		deleted, err := r.RouteRepo.Delete(ctx, tx, route)
		return utils.SafeReturn(&deleted, nil, err)
	})
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
