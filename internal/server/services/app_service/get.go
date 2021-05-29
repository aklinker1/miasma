package app_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Get(tx *gorm.DB, appName string) (*models.App, error) {
	log.V("app_service.Get(%v)", appName)
	var app models.App
	err := tx.First(&app, "name = ?", appName).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}
