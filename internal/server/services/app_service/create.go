package app_service

import (
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"github.com/dgryski/trifles/uuid"
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func Create(tx *gorm.DB, app models.AppInput) (*models.App, error) {
	log.V("app_service.Create(%v)", app)
	newApp := &models.App{
		ID:     strfmt.UUID4(uuid.UUIDv4()),
		Name:   app.Name,
		Group:  app.Group,
		Image:  app.Image,
		Hidden: app.Hidden,
	}
	err := tx.Create(newApp).Error
	if err != nil {
		return nil, err
	}

	return newApp, nil
}
