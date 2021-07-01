package app_service

import (
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Delete(tx *gorm.DB, app *models.App) error {
	log.V("app_service.Delete(%v)", app)
	err := docker_service.StopService(app.Name)
	if err != nil {
		log.V("No need to stop %s (%v)", app.Name, err)
	} else {
		log.V("Stopped %s before deleting", app.Name)
	}

	err = docker_service.DeleteNetwork(app.Name)
	if err != nil {
		log.W("Failed to destroy network: %v", err)
	}

	return tx.Delete(&app).Error
}
