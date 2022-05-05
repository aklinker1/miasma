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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditApp(ctx context.Context, appName string, app internal.AppChanges) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteApp(ctx context.Context, appName string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StartApp(ctx context.Context, appName string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StopApp(ctx context.Context, appName string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReloadApp(ctx context.Context, appName string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpgradeApp(ctx context.Context, appName string) (*internal.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EnablePlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DisablePlugin(ctx context.Context, pluginName string) (*internal.Plugin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SetAppRouting(ctx context.Context, appName string, routing *internal.AppRoutingInput) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveAppRouting(ctx context.Context, appName string) (*internal.AppRouting, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
