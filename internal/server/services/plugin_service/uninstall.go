package plugin_service

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Uninstall(tx *gorm.DB, pluginName string) (plugin *models.Plugin, pluginApp *models.App, err error) {
	switch pluginName {
	case constants.PluginNameTraefik:
		plugin, pluginApp, err = uninstallTraefik(tx)
	default:
		err = fmt.Errorf("%s is not a valid plugin name", pluginName)
	}

	return plugin, pluginApp, err
}

func uninstallTraefik(tx *gorm.DB) (*models.Plugin, *models.App, error) {
	if !IsInstalled(tx, constants.PluginNameTraefik) {
		return nil, nil, fmt.Errorf("traefik is not installed")
	}

	err := UpdatePluginInstalled(tx, constants.PluginNameTraefik, false)
	if err != nil {
		return nil, nil, err
	}

	plugin, err := Get(tx, constants.PluginNameTraefik)
	return plugin, constants.Plugins.Traefik.App, err
}
