package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server/gqlgen"
)

func (r *mutationResolver) CreateApp(ctx context.Context, app internal.AppInput) (*internal.App, error) {
	created, err := r.Apps.Create(ctx, app)
	return safeReturn(&created, nil, err)
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
	panic(fmt.Errorf("not implemented"))
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
