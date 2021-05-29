package plugin_service

import (
	"fmt"

	"github.com/aklinker1/miasma/package/models"
)

func Uninstall(pluginName string) (plugin *models.Plugin, err error) {
	switch pluginName {
	case "traefik":
		plugin, err = uninstallTraefik()
	default:
		err = fmt.Errorf("%s is not a valid plugin name", pluginName)
	}

	return plugin, err
}

func uninstallTraefik() (*models.Plugin, error) {
	panic("TODO: Save that the app is uninstalled")
	// pluginMeta.Traefik = false
	// traefik := constants.Plugins.Traefik

	// err := app_service.StopApp(traefik.Name)
	// if err != nil {
	// 	return nil, err
	// }
	// err = docker_service.DestroyNetwork(traefik.Name)
	// if err != nil {
	// 	log.W("Failed to destroy network: %v", err)
	// }
	// _ = WritePluginMeta(pluginMeta)
	// if err != nil {
	// 	log.W("Failed to mark Traefik as uninstalled: %v", err)
	// }

	panic("TODO: remove all the apps plugins for traefik")

	// return Get(traefik.Name, pluginMeta)
}
