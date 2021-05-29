package env_service

import (
	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/internal/shared/log"
	"gorm.io/gorm"
)

func Get(tx *gorm.DB, appID string) (map[string]string, error) {
	log.V("env_service.Get(%v)", appID)
	envVars := []entities.SQLEnvVar{}
	err := tx.Where("app_id = ?", appID).Find(&envVars).Error

	keyValues := map[string]string{}
	if err != nil {
		return keyValues, err
	}
	for _, envVar := range envVars {
		keyValues[envVar.Key] = envVar.Value
	}

	return keyValues, nil
}
