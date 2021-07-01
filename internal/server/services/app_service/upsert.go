package app_service

import (
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Upsert(tx *gorm.DB, app *models.App) error {
	return tx.Save(app).Error
}
