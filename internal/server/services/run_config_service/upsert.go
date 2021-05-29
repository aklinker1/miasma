package run_config_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Upsert(tx *gorm.DB, newRunConfig *models.RunConfig) error {
	log.V("Updating config for %s...", newRunConfig.AppID)
	panic("TODO: save new run config")
	// existingRunConfig, err := app_service.GetAppMeta(appName)
	// if err != nil {
	// 	return nil, err
	// }

	// TODO! Reset ports to be generated if the published ports have changed
	panic("TODO: Pull out reload into parent function")
	// err = app_service.Reload(appName, updatedMeta)
	// if err != nil {
	// 	return nil, err
	// }
}
