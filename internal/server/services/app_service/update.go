package app_service

import (
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Update(tx *gorm.DB, newApp *models.App) error {
	return tx.Save(newApp).Error
}
