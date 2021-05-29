package app_service

import (
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/services/run_config_service"
	"github.com/aklinker1/miasma/internal/server/utils/mappers/app_mapper"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func Reload(details *server_models.AppDetails, env map[string]string, plugins *server_models.AppPlugins) error {
	log.V("app_service.reload(%v, env:***, %v)", details, plugins)
	newServiceSpec, err := app_mapper.ToService(details, env, plugins, run_config_service.GetNextPorts(details), details.RunConfig.ImageDigest)
	if err != nil {
		return err
	}

	existingService, _ := docker_service.GetRunningService(details.App.Name)
	if existingService == nil {
		log.V("%s is not running, do not need to update it", details.App.Name)
		return nil
	}

	log.V("Updating a running service: %s", details.App.Name)
	return docker_service.UpdateService(existingService, newServiceSpec)
}
