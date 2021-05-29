package run_config_service

import (
	"github.com/aklinker1/miasma/internal/server/utils/mappers/run_config_mapper"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Create(tx *gorm.DB, runConfig *models.RunConfig) error {
	entity := run_config_mapper.ToSQL(runConfig)
	err := tx.Create(entity).Error
	if err != nil {
		return err
	}
	return nil
}
