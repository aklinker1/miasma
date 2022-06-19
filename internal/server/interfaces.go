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
	// Create and start a new app
	Create(ctx context.Context, app internal.App) (internal.App, error)
	// FindApps searches the list of managed applications
	FindApps(ctx context.Context, filter AppsFilter) ([]internal.App, error)
	// FindApp searches the list of managed applications for the first app that matches the criteria
	FindApp(ctx context.Context, filter AppsFilter) (internal.App, error)
	// Update all the fields on an app and restart it. If a new image is passed, update the image digest.
	Update(ctx context.Context, app internal.App, newImage *string) (internal.App, error)
	// Delete and stop the app
	Delete(ctx context.Context, name string) (internal.App, error)
}

// RuntimeService defines how the server runs the apps
type RuntimeService interface {
	// Start the app
	Start(ctx context.Context, app internal.App) error
	// Restart stops and starts the app
	Restart(ctx context.Context, app internal.App) error
	// Stop stops the app if it's running
	Stop(ctx context.Context, app internal.App) error
	// PullLatest grabs the latest image and returns it's digest
	PullLatest(ctx context.Context, image string) (string, error)
	// Version returns the runtime's version
	Version(ctx context.Context) (string, error)
	// ClusterInfo returns details about the device cluster
	ClusterInfo(ctx context.Context) (*internal.ClusterInfo, error)
}

type RoutesFilter struct {
	AppID *string
}

type RouteService interface {
	// FindRoute returns the first route matching the filter
	FindRoute(ctx context.Context, filter RoutesFilter) (internal.AppRouting, error)
}
