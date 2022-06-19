package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/gofrs/uuid"
)

func (r *mutationResolver) CreateApp(ctx context.Context, app internal.AppInput) (*internal.App, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to generate new app's ID",
			Op:      "sqlite.createApp",
			Err:     err,
		}
	}
	if app.Name = strings.TrimSpace(app.Name); app.Name == "" {
		return nil, &server.Error{
			Code:    server.EINVALID,
			Message: "App name cannot be empty",
			Op:      "createApp",
		}
	}
	if app.Image = strings.TrimSpace(app.Image); app.Image == "" {
		return nil, &server.Error{
			Code:    server.EINVALID,
			Message: "App image cannot be empty",
			Op:      "createApp",
		}
	}

	a := internal.App{
		ID:             id.String(),
		CreatedAt:      time.Now(),
		Name:           app.Name,
		Group:          app.Group,
		Image:          app.Image,
		ImageDigest:    "TODO",
		Hidden:         utils.BoolOr(app.Hidden, false),
		TargetPorts:    app.TargetPorts,
		PublishedPorts: app.PublishedPorts,
		Placement:      app.Placement,
		Volumes:        toBoundVolumes(app.Volumes),
		Networks:       app.Networks,
		Command:        app.Command,
	}

	return safeReturn(&a, nil, r.Apps.Create(ctx, a))
}

func (r *mutationResolver) EditApp(ctx context.Context, id string, changes map[string]interface{}) (*internal.App, error) {
	newApp, err := r.Apps.GetOne(ctx, internal.AppsFilter{ID: &id})
	if err != nil {
		return nil, err
	}
	gqlgen.ApplyChanges(changes, &newApp)
	updated, err := r.Apps.Update(ctx, newApp)
	return safeReturn(&updated, nil, err)
}

func (r *mutationResolver) DeleteApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.Delete(ctx, appName)
	return safeReturn(&app, nil, err)
}

func (r *mutationResolver) StartApp(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StopApp(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReloadApp(ctx context.Context, id string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpgradeApp(ctx context.Context, id string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EnablePlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DisablePlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetAppRouting(ctx context.Context, appID string, routing *internal.AppRoutingInput) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveAppRouting(ctx context.Context, appID string) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
