package run_config_service

import (
	"github.com/aklinker1/miasma/internal/server/utils/mappers/run_config_mapper"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Upsert(tx *gorm.DB, newRunConfig *models.RunConfig) error {
	log.V("run_config_service.Upsert(%v)", newRunConfig)

	entity := run_config_mapper.ToSQL(newRunConfig)
	return tx.Save(entity).Error
}
