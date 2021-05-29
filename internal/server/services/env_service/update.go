package env_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"gorm.io/gorm"
)

func Update(tx *gorm.DB, appId string, newEnv map[string]interface{}) error {
	log.V("env_service.Update(%v, env:***)", appId)
	panic("TODO: Pull app reload out of env_service")
	// err = app_service.Reload(appName, updatedMeta)
	// if err != nil {
	// 	return nil, err
	// }
}
