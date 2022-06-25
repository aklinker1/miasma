package cli

import (
	"context"

	"github.com/aklinker1/miasma/internal"
)

type Metadata struct {
	Version   string
	Build     string
	BuildHash string
	BuildDate string
}

type ListAppsOptions struct {
	ShowHidden *bool `json:"showHidden"`
}

type APIService interface {
	SetBaseURL(baseURL string)

	Health(ctx context.Context, gql string) (internal.Health, error)

	ListApps(ctx context.Context, options ListAppsOptions, gql string) ([]internal.App, error)
	GetApp(ctx context.Context, appName string, gql string) (internal.App, error)
	CreateApp(ctx context.Context, app internal.AppInput) error
	DeleteApp(ctx context.Context, appID string) error
	StartApp(ctx context.Context, appID string) error
	StopApp(ctx context.Context, appID string) error
	RestartApp(ctx context.Context, appID string) error

	SetAppEnv(ctx context.Context, env internal.EnvMap) error

	SetAppRoute(ctx context.Context, appID string, route internal.RouteInput) error
	RemoveAppRoute(ctx context.Context, appID string) error

	ListPlugins(ctx context.Context, gql string) ([]internal.Plugin, error)
	EnablePlugin(ctx context.Context, pluginName string) error
	DisablePlugin(ctx context.Context, pluginName string) error
}
