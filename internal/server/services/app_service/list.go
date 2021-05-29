package app_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func List(tx *gorm.DB, showHidden bool) ([]*models.App, error) {
	log.V("app_service.List(%v)", showHidden)
	apps := []*models.App{}
	query := tx.Order("`group` ASC, name ASC")
	if !showHidden {
		query = query.Where("hidden = ?", false)
	}
	err := query.Find(&apps).Error
	if err != nil {
		return nil, err
	}
	return apps, nil
}
