package app_service

import (
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
)

func Stop(app *models.App) error {
	log.V("app_service.Start(%v)", app)
	return docker_service.StopService(app.Name)
}
