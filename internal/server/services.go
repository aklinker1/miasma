package server

import (
	"context"

	"github.com/aklinker1/miasma/internal"
)

type AppService interface {
	Create(ctx context.Context, app internal.AppInput) (internal.App, error)
	Get(ctx context.Context, filter internal.AppsFilter) ([]internal.App, error)
	GetOne(ctx context.Context, filter internal.AppsFilter) (internal.App, error)
	Update(ctx context.Context, app internal.App) (internal.App, error)
	Delete(ctx context.Context, appName string) (internal.App, error)
}

// RuntimeService defines how the server runs the apps
type RuntimeService interface {
	Start(ctx context.Context, app internal.App) error
	Restart(ctx context.Context, app internal.App) error
	Stop(ctx context.Context, app internal.App) error
	PullLatest(ctx context.Context, image string) (string, error)
	Version(ctx context.Context) (string, error)
	SwarmInfo(ctx context.Context) (*internal.SwarmInfo, error)
}
