package plugin_service

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Get(tx *gorm.DB, pluginName string) (*models.Plugin, error) {
	log.V("plugin_server.Get(%v)", pluginName)

	installed := false
	found := false
	switch pluginName {
	// case "postgres":
	case constants.PluginNameTraefik:
		installed = IsInstalled(tx, pluginName)
		found = true
	}
	if !found {
		return nil, fmt.Errorf("No plugin named %v is available", pluginName)
	}
	return &models.Plugin{
		Name:      pluginName,
		Installed: installed,
	}, nil
}
