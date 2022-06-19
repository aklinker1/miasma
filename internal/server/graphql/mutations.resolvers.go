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
	"github.com/samber/lo"
)

func (r *mutationResolver) CreateApp(ctx context.Context, input internal.AppInput) (*internal.App, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to generate new app's ID",
			Op:      "sqlite.createApp",
			Err:     err,
		}
	}
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
		ID:             id.String(),
		CreatedAt:      time.Now(),
		Name:           input.Name,
		Group:          input.Group,
		Image:          input.Image,
		ImageDigest:    "TODO",
		Hidden:         utils.BoolOr(input.Hidden, false),
		Routing:        (*internal.AppRouting)(input.Routing),
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
	updated, err := r.Apps.Update(ctx, newApp)
	return safeReturn(&updated, nil, err)
}

func (r *mutationResolver) DeleteApp(ctx context.Context, id string) (*internal.App, error) {
	app, err := r.Apps.Delete(ctx, id)
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

func (r *mutationResolver) EnablePlugin(ctx context.Context, name string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DisablePlugin(ctx context.Context, name string) (*internal.Plugin, error) {
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
