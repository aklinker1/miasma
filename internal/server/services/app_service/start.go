package app_service

import (
	"fmt"

	"docker.io/go-docker/api/types"
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/services/env_service"
	"github.com/aklinker1/miasma/internal/server/services/run_config_service"
	"github.com/aklinker1/miasma/internal/server/utils/mappers/app_mapper"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
	"gorm.io/gorm"
)

func Start(tx *gorm.DB, details *server_models.AppDetails, plugins *server_models.AppPlugins) error {
	log.V("app_service.Start(%v, %v)", details, plugins)

	existingService, _ := docker_service.GetRunningService(details.App.Name)
	if existingService != nil {
		return fmt.Errorf("%s is already running", details.App.Name)
	}

	env, err := env_service.Get(tx, details.App.ID.String())
	if err != nil {
		return err
	}

	newService, err := app_mapper.ToService(
		details,
		env,
		plugins,
		run_config_service.GetNextPorts(details),
		details.RunConfig.ImageDigest,
	)
	if err != nil {
		return err
	}

	// Ensure the app's network exists
	err = docker_service.UpsertNetwork(details.App.Name)
	if err != nil {
		return err
	}

	options := types.ServiceCreateOptions{
		QueryRegistry: true,
	}
	return docker_service.StartService(*newService, options)
}
