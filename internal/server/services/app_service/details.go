package app_service

import (
	"github.com/aklinker1/miasma/internal/server/services/run_config_service"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
	"gorm.io/gorm"
)

func Details(tx *gorm.DB, appName string) (*server_models.AppDetails, error) {
	log.V("app_service.Details(%v)", appName)
	app, err := Get(tx, appName)
	if err != nil {
		return nil, err
	}
	appId := app.ID.String()

	runConfig, err := run_config_service.Get(tx, appId)
	if err != nil {
		return nil, err
	}

	return &server_models.AppDetails{
		App:       app,
		RunConfig: runConfig,
	}, nil
}
