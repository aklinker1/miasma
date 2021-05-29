package plugin_service

import (
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func UpsertTraefikConfig(tx *gorm.DB, config *models.TraefikPluginConfig) error {
	return tx.Save(config).Error
}
