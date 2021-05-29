package plugin_service

import (
	"errors"
	"fmt"

	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func IsInstalled(tx *gorm.DB, pluginName string) bool {
	log.V("plugin_service.IsInstalled(%v)", pluginName)
	var plugin models.Plugin
	err := tx.Find(&plugin, "name = ?", pluginName).Error
	if err == nil {
		return plugin.Installed
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	panic(fmt.Sprintf("Unknown error when checking if a plugin is installed (%s)", err.Error()))
}
