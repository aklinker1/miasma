package services

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
)

type PluginService struct {
	DB               server.DB
	Logger           server.Logger
	PluginRepo       server.PluginRepo
	CertResolverName string
	RuntimeService   *RuntimeService
}

var (
	pluginAppGroup = lo.ToPtr("System")
)

func (s *PluginService) TogglePlugin(ctx context.Context, enabled bool, name internal.PluginName, config map[string]any) (internal.Plugin, error) {
	return utils.InTx(ctx, s.DB.ReadWriteTx, zero.Plugin, func(tx server.Tx) (internal.Plugin, error) {
		plugin, err := s.PluginRepo.GetOne(ctx, tx, server.PluginsFilter{
			Name: &name,
		})
		if err != nil {
			return zero.Plugin, err
		}

		// Stop if no change
		if plugin.Enabled == enabled {
			s.Logger.W("Plugin already %s", lo.Ternary(enabled, "enabled", "disabled"))
			return plugin, nil
		}

		// Update the plugin
		plugin.Enabled = enabled
		plugin.Config = config
		updated, err := s.PluginRepo.Update(ctx, tx, plugin)
		if err != nil {
			return zero.Plugin, err
		}

		// Execute the enabled/disabled hooks for each plugin
		if enabled {
			err = s.onPluginEnabled(ctx, tx, updated)
		} else {
			err = s.onPluginDisabled(ctx, tx, updated)
		}
		return updated, err
	})
}

func (s *PluginService) onPluginEnabled(ctx context.Context, tx server.Tx, plugin internal.Plugin) error {
	switch plugin.Name {
	case internal.PluginNameTraefik:
		return s.onTraefikEnabled(ctx, tx, plugin)
	}
	return nil
}

func (s *PluginService) onTraefikEnabled(ctx context.Context, tx server.Tx, plugin internal.Plugin) error {
	// Start Traefik
	app, err := s.createTraefikApp(plugin.ConfigForTraefik())
	if err != nil {
		return err
	}
	err = s.RuntimeService.StartApp(ctx, tx, PartialRuntimeServiceSpec{
		App: app,
	})
	if err != nil {
		return err
	}

	// Restart all other apps
	return s.RuntimeService.RestartAllAppsIfRunning(ctx, tx)
}

func (s *PluginService) onPluginDisabled(ctx context.Context, tx server.Tx, plugin internal.Plugin) error {
	switch plugin.Name {
	case internal.PluginNameTraefik:
		return s.onTraefikDisabled(ctx, tx, plugin)
	}
	return nil
}

func (s *PluginService) onTraefikDisabled(ctx context.Context, tx server.Tx, plugin internal.Plugin) error {
	// Stop Traefik
	app, err := s.createTraefikApp(plugin.ConfigForTraefik())
	if err != nil {
		return err
	}
	err = s.RuntimeService.StopApp(ctx, tx, PartialRuntimeServiceSpec{
		App: app,
	})
	if err != nil {
		return err
	}

	// Restart all other apps
	return s.RuntimeService.RestartAllAppsIfRunning(ctx, tx)
}

func (s *PluginService) createTraefikApp(config internal.TraefikConfig) (internal.App, error) {
	command := []string{"traefik"}
	if config.EnableHttps {
		if config.CertsEmail == "" {
			return zero.App, &server.Error{
				Code:    server.EINVALID,
				Message: "Certificate email is missing, did you provide \"certsEmail\" in the config?",
				Op:      "sqlite.PluginService.traefikApp",
			}
		}
		command = append(
			command,
			"--entrypoints.web.address=:80",
			"--entrypoints.websecure.address=:443",
			// Use LetsEncrypt to manage certs: https://doc.traefik.io/traefik/https/acme/#configuration-examples
			fmt.Sprintf("--certificatesresolvers.%s.acme.email=%s", s.CertResolverName, config.CertsEmail),
			fmt.Sprintf("--certificatesresolvers.%s.acme.storage=/letsencrypt/acme.json", s.CertResolverName),
			fmt.Sprintf("--certificatesresolvers.%s.acme.httpchallenge.entrypoint=web", s.CertResolverName),
			// Redirect HTTP -> HTTPS: https://doc.traefik.io/traefik/routing/entrypoints/#redirection
			"--entrypoints.web.http.redirections.entrypoint.to=websecure",
			"--entrypoints.web.http.redirections.entrypoint.scheme=https",
		)
	}
	command = append(command, "--api.insecure=true", "--providers.docker", "--providers.docker.swarmmode")

	ports := []int{80}
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
		Group:          pluginAppGroup,
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
