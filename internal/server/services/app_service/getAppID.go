package app_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"gorm.io/gorm"
)

func GetAppID(tx *gorm.DB, appName string) (string, error) {
	log.V("app_service.GetAppID(%v)", appName)
	app, err := Get(tx, appName)
	if err != nil {
		return "", err
	}
	return app.ID.String(), nil
}
