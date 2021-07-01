package env_service

import (
	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func Update(tx *gorm.DB, appID string, newEnv map[string]interface{}) error {
	log.V("env_service.Update(%v, env:***)", appID)

	// Remove old variables
	existingVars, err := Get(tx, appID)
	if err != nil {
		return err
	}
	for key := range existingVars {
		_, ok := newEnv[key]
		if !ok {
			err := tx.Delete(&entities.SQLEnvVar{}, "app_id = ? AND key = ?", appID, key).Error
			if err != nil {
				return err
			}
		}
	}

	// Add new variables
	for key, value := range newEnv {
		err := tx.Save(&entities.SQLEnvVar{
			AppID: strfmt.UUID4(appID),
			Key:   key,
			Value: value.(string),
		}).Error
		if err != nil {
			return err
		}
	}

	return nil
}
