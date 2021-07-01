package run_config_service

import (
	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/internal/server/utils/mappers/run_config_mapper"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Get(tx *gorm.DB, appID string) (*models.RunConfig, error) {
	var runConfig entities.SQLRunConfig
	err := tx.First(&runConfig, "app_id = ?", appID).Error
	if err != nil {
		return nil, err
	}
	return run_config_mapper.FromSQL(&runConfig), nil
}
