package sqlite

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/samber/lo"
)

var (
	EmptyPlugin    = internal.Plugin{}
	PluginAppGroup = lo.ToPtr("System")
)

type PluginService struct {
	db               server.DB
	logger           server.Logger
	apps             server.AppService
	runtime          server.RuntimeService
	dataDir          string
	certResolverName string
}

func NewPluginService(db server.DB, apps server.AppService, runtime server.RuntimeService, logger server.Logger, certResolverName string) server.PluginService {
	return &PluginService{
		db:               db,
		logger:           logger,
		apps:             apps,
		runtime:          runtime,
		certResolverName: certResolverName,
	}
}

func (s *PluginService) getTraefikApp(config internal.TraefikConfig) (internal.App, error) {
	s.logger.V("Traefik app config: %+v", config)

	command := []string{"traefik"}
	if config.EnableHttps {
		if config.CertEmail == "" {
			return EmptyApp, &server.Error{
				Code:    server.EINVALID,
				Message: "Certificate email is missing, did you provide \"certEmail\" in the config?",
				Op:      "sqlite.PluginService.traefikApp",
			}
		}
		command = append(
			command,
			"--entrypoints.web.address=:80",
			"--entrypoints.websecure.address=:443",
			// Use LetsEncrypt to manage certs: https://doc.traefik.io/traefik/https/acme/#configuration-examples
			fmt.Sprintf("--certificatesresolvers.%s.acme.email=%s", s.certResolverName, config.CertEmail),
			fmt.Sprintf("--certificatesresolvers.%s.acme.storage=/letsencrypt/acme.json", s.certResolverName),
			fmt.Sprintf("--certificatesresolvers.%s.acme.httpchallenge.entrypoint=web", s.certResolverName),
			// Redirect HTTP -> HTTPS: https://doc.traefik.io/traefik/routing/entrypoints/#redirection
			"--entrypoints.web.http.redirections.entrypoint.to=websecure",
			"--entrypoints.web.http.redirections.entrypoint.scheme=https",
		)
	}
	command = append(command, "--api.insecure=true", "--providers.docker", "--providers.docker.swarmmode")

	ports := []int32{80}
	if config.EnableHttps {
		ports = append(ports, 443)
	}
	ports = append(ports, 8080)

	volumes := []*internal.BoundVolume{{
		Source: "/var/run/docker.sock",
		Target: "/var/run/docker.sock",
	}}
	if config.EnableHttps {
		volumes = append(volumes, &internal.BoundVolume{
			Source: config.CertsDir,
			Target: "/letsencrypt",
		})
	}

	return internal.App{
		ID:             "plugin-traefik",
		Name:           "Traefik",
		Group:          PluginAppGroup,
		System:         true,
		Hidden:         true,
		Image:          "traefik:2.7",
		ImageDigest:    "sha256:fdff55caa91ac7ff217ff03b93f3673844b3b88ad993e023ab43f6004021697c",
		TargetPorts:    ports,
		PublishedPorts: ports,
		Placement:      []string{"node.role == manager"},
		Volumes:        volumes,
		Command:        command,
	}, nil
}

func (s *PluginService) setEnabled(ctx context.Context, plugin internal.Plugin, enabled bool) (internal.Plugin, error) {
	if plugin.Enabled == enabled {
		s.logger.W("Plugin already %s", lo.Ternary(enabled, "enabled", "disabled"))
		return plugin, nil
	}

	tx, err := s.db.ReadWriteTx(ctx)
	if err != nil {
		return EmptyPlugin, err
	}
	defer tx.Rollback()

	plugin.Enabled = enabled
	updated, err := updatePlugin(ctx, tx, plugin)
	if err != nil {
		return EmptyPlugin, err
	}

	// Execute setup/teardown
	if enabled {
		err = s.onEnabled(ctx, tx, updated)
	} else {
		err = s.onDisabled(ctx, tx, updated)
	}
	if err != nil {
		return EmptyPlugin, err
	}

	// Restart managed apps
	apps, err := findApps(ctx, tx, server.AppsFilter{})
	if err != nil {
		return EmptyPlugin, err
	}
	params := []server.StartAppParams{}
	for _, app := range apps {
		route, err := findRouteOrNil(ctx, tx, server.RoutesFilter{
			AppID: &app.ID,
		})
		if err != nil {
			return EmptyPlugin, err
		}
		env, err := findEnvMap(ctx, tx, server.EnvFilter{
			AppID: &app.ID,
		})
		if err != nil {
			return EmptyPlugin, err
		}
		params = append(params, server.StartAppParams{
			App:   app,
			Route: route,
			Env:   env,
		})
	}
	s.runtime.RestartRunningApps(ctx, params)
	if err != nil {
		return EmptyPlugin, err
	}

	tx.Commit()
	return updated, nil
}

// DisablePlugin implements server.PluginService
func (s *PluginService) DisablePlugin(ctx context.Context, plugin internal.Plugin) (internal.Plugin, error) {
	return s.setEnabled(ctx, plugin, false)
}

// EnablePlugin implements server.PluginService
func (s *PluginService) EnablePlugin(ctx context.Context, plugin internal.Plugin, config map[string]any) (internal.Plugin, error) {
	plugin.Config = config
	return s.setEnabled(ctx, plugin, true)
}

// FindPlugins implements server.PluginService
func (s *PluginService) FindPlugins(ctx context.Context, filter server.PluginsFilter) ([]internal.Plugin, error) {
	s.logger.D("Finding plugins that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return nil, err
	}
	return findPlugins(ctx, tx, filter)
}

// FindPlugin implements server.PluginService
func (s *PluginService) FindPlugin(ctx context.Context, filter server.PluginsFilter) (internal.Plugin, error) {
	s.logger.D("Finding plugins that matches: %+v", filter)
	tx, err := s.db.ReadonlyTx(ctx)
	if err != nil {
		return EmptyPlugin, err
	}
	return findPlugin(ctx, tx, filter)
}

func (s *PluginService) onEnabled(ctx context.Context, tx server.Tx, plugin internal.Plugin) error {
	s.logger.D("On plugin enabled: %v", plugin.Name)
	switch plugin.Name {
	case internal.PluginNameTraefik:
		app, err := s.getTraefikApp(plugin.ConfigForTraefik())
		if err != nil {
			return err
		}
		created, err := createApp(ctx, tx, app)
		if err != nil {
			return err
		}
		route, err := findRouteOrNil(ctx, tx, server.RoutesFilter{
			AppID: &created.ID,
		})
		if err != nil {
			return err
		}
		env, err := findEnvMap(ctx, tx, server.EnvFilter{
			AppID: &created.ID,
		})
		if err != nil {
			return err
		}
		allPlugins, err := findPlugins(ctx, tx, server.PluginsFilter{})
		if err != nil {
			return err
		}
		return s.runtime.Start(ctx, created, route, env, allPlugins)
	default:
		s.logger.V("No onEnabled hook for %v", plugin.Name)
	}
	return nil
}

func (s *PluginService) onDisabled(ctx context.Context, tx server.Tx, plugin internal.Plugin) error {
	s.logger.D("On plugin disabled: %v", plugin)
	switch plugin.Name {
	case internal.PluginNameTraefik:
		app, err := s.getTraefikApp(plugin.ConfigForTraefik())
		if err != nil {
			return err
		}
		err = deleteApp(ctx, tx, app)
		if err != nil {
			return err
		}
		return s.runtime.Stop(ctx, app)
	default:
		s.logger.V("No onDisabled hook for %v", plugin)
	}
	return nil
}
