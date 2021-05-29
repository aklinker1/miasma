package plugin_service

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/package/models"
)

func Get(pluginName string) (*models.Plugin, error) {
	panic("TODO: Lookup installed plugins")
	installed := false
	switch pluginName {
	case "traefik":
		panic("TODO: Check if traefik is installed")
		// installed = meta.Traefik
		// case "postgres":
		// 	installed = meta.Postgres
		// case "mongo":
		// 	installed = meta.Mongo
	}

	var installCommand *string
	var uninstallCommand *string
	if installed {
		uninstallCommand = shared.StringPtr(fmt.Sprintf("miasma plugin:uninstall %s", pluginName))
	} else {
		installCommand = shared.StringPtr(fmt.Sprintf("miasma plugin:install %s", pluginName))
	}

	return &models.Plugin{
		Name:             &pluginName,
		Installed:        &installed,
		InstallCommand:   installCommand,
		UninstallCommand: uninstallCommand,
	}, nil
}
