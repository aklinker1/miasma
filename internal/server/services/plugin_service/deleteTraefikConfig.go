package plugin_service

import (
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func DeleteTraefikConfig(tx *gorm.DB, appID string) error {
	return tx.Delete(&models.TraefikPluginConfig{}, "app_id = ?", appID).Error
}
