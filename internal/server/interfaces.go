package server

import (
	"context"
	"database/sql"

	"github.com/aklinker1/miasma/internal"
)

type Tx = *sql.Tx

type DB interface {
	Open() error
	ReadonlyTx(ctx context.Context) (Tx, error)
	ReadWriteTx(ctx context.Context) (Tx, error)
}

type Server interface {
	ServeGraphql() error
}

type Logger interface {
	D(format string, args ...any)
	V(format string, args ...any)
	I(format string, args ...any)
	W(format string, args ...any)
	E(format string, args ...any)
}

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
	Create(ctx context.Context, app internal.App, plugins []internal.Plugin) (internal.App, error)
	// FindApps searches the list of managed applications
	FindApps(ctx context.Context, filter AppsFilter) ([]internal.App, error)
	// FindApp searches the list of managed applications for the first app that matches the criteria
	FindApp(ctx context.Context, filter AppsFilter) (internal.App, error)
	// Update all the fields on an app and restart it. If a new image is passed, update the image digest.
	Update(ctx context.Context, app internal.App, newImage *string) (internal.App, error)
	// Delete and stop the app
	Delete(ctx context.Context, name string) (internal.App, error)
}

type PluginsFilter struct {
	Name         *internal.PluginName
	NameContains *string
	Enabled      *bool
}

type PluginService interface {
	// FindPlugins searches the list of built-in plugins
	FindPlugins(ctx context.Context, filter PluginsFilter) ([]internal.Plugin, error)
	// FindApp searches the list of built-in plugins for the first plugin that matches the criteria
	FindPlugin(ctx context.Context, filter PluginsFilter) (internal.Plugin, error)
	// Enabled a plugin and restart all applications
	EnablePlugin(ctx context.Context, plugin internal.Plugin, config map[string]any) (internal.Plugin, error)
	// Disable a plugin and restart all applications
	DisablePlugin(ctx context.Context, plugin internal.Plugin) (internal.Plugin, error)
}

type RuntimeAppInfo struct {
	Instances      internal.AppInstances
	Status         string
	PublishedPorts []uint32
}

type StartAppParams struct {
	App     internal.App
	Route   *internal.Route
	Env     internal.EnvMap
	Plugins []internal.Plugin
}

type ListServicesFilter struct {
	NodeID *string
}

// RuntimeService defines how the server runs the apps
type RuntimeService interface {
	// Start the app
	Start(ctx context.Context, app internal.App, route *internal.Route, env map[string]string, plugins []internal.Plugin) error
	// ServiceDetails returns runtime details like instance count and status
	GetRuntimeAppInfo(ctx context.Context, app internal.App) (RuntimeAppInfo, error)
	// Restart stops and starts the app
	Restart(ctx context.Context, app internal.App, route *internal.Route, env map[string]string, plugins []internal.Plugin) error
	// Stop stops the app if it's running
	Stop(ctx context.Context, app internal.App) error
	// PullLatest grabs the latest image and returns it's digest
	PullLatest(ctx context.Context, image string) (string, error)
	// Version returns the runtime's version
	Version(ctx context.Context) (string, error)
	// ClusterInfo returns details about the device cluster
	ClusterInfo(ctx context.Context) (*internal.ClusterInfo, error)
	RestartRunningApps(ctx context.Context, params []StartAppParams) error
	ListNodes(ctx context.Context) ([]internal.Node, error)
	ListServices(ctx context.Context, filter ListServicesFilter) ([]internal.RunningContainer, error)
}

type RoutesFilter struct {
	AppID *string
}

type RouteService interface {
	// FindRoute returns the first route matching the filter
	FindRoute(ctx context.Context, filter RoutesFilter) (internal.Route, error)
	FindRouteOrNil(ctx context.Context, filter RoutesFilter) (*internal.Route, error)
	Create(ctx context.Context, route internal.Route) (internal.Route, error)
	Update(ctx context.Context, route internal.Route) (internal.Route, error)
	Delete(ctx context.Context, route internal.Route) (internal.Route, error)
}

type EnvFilter struct {
	AppID *string
}

type EnvService interface {
	FindEnv(ctx context.Context, filter EnvFilter) (internal.EnvMap, error)
	SetAppEnv(ctx context.Context, appID string, newEnv internal.EnvMap) (internal.EnvMap, error)
}
