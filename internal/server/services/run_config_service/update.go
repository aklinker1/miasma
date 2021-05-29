package run_config_service

import (
	"github.com/aklinker1/miasma/internal/server/utils/mappers/run_config_mapper"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Update(tx *gorm.DB, newRunConfig *models.RunConfig) error {
	entity := run_config_mapper.ToSQL(newRunConfig)
	return tx.Save(entity).Error
}
