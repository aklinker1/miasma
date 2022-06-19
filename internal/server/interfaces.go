package server

import (
	"context"

	"github.com/aklinker1/miasma/internal"
)

type Pagination struct {
	Page int32
	Size int32
}

func (p Pagination) Limit() int32 {
	return p.Size
}

func (p Pagination) Offset() int32 {
	zeroIndexPage := p.Page - 1
	if zeroIndexPage < 0 {
		zeroIndexPage = 0
	}
	return zeroIndexPage * p.Size
}

type Sort struct {
	Field     string
	Direction string
}

type AppsFilter struct {
	ID            *string
	Name          *string
	NameContains  *string
	IncludeHidden *bool
	Sort          *Sort
	Pagination    *Pagination
}

type AppService interface {
	Create(ctx context.Context, app internal.App) (internal.App, error)
	FindApps(ctx context.Context, filter AppsFilter) ([]internal.App, error)
	FindApp(ctx context.Context, filter AppsFilter) (internal.App, error)
	Update(ctx context.Context, app internal.App) (internal.App, error)
	Delete(ctx context.Context, name string) (internal.App, error)
}

// RuntimeService defines how the server runs the apps
type RuntimeService interface {
	Start(ctx context.Context, app internal.App) error
	Restart(ctx context.Context, app internal.App) error
	Stop(ctx context.Context, app internal.App) error
	PullLatest(ctx context.Context, image string) (string, error)
	Version(ctx context.Context) (string, error)
	ClusterInfo(ctx context.Context) (*internal.ClusterInfo, error)
}

type RoutesFilter struct {
	AppID *string
}

type RouteService interface {
	FindRoute(ctx context.Context, filter RoutesFilter) (internal.AppRouting, error)
}
