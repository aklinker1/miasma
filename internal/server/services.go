package server

import (
	"context"

	"github.com/aklinker1/miasma/internal"
)

type GetAppFilters struct {
	Pagination
	Name          string
	IncludeHidden bool
}

type AppService interface {
	Create(ctx context.Context, app internal.App) error
	FindAppByName(ctx context.Context, name string) (internal.App, error)
	FindApps(ctx context.Context, filter GetAppFilters) ([]internal.App, error)
	Update(ctx context.Context, app internal.App) error
	Delete(ctx context.Context, name string) (internal.App, error)
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
