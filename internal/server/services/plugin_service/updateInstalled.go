package plugin_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func UpdatePluginInstalled(tx *gorm.DB, pluginName string, installed bool) error {
	log.V("plugin_service.UpdatePluginInstalled(%v, %v)", pluginName, installed)
	return tx.Save(&models.Plugin{
		Name:      pluginName,
		Installed: installed,
	}).Error
}
