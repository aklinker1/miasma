package plugin_service

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func Install(tx *gorm.DB, pluginName string) (plugin *models.Plugin, pluginDetails *server_models.AppDetails, err error) {
	log.V("plugin_server.Install(%v)", pluginName)

	switch pluginName {
	case constants.PluginNameTraefik:
		plugin, pluginDetails, err = installTraefik(tx)
	// case "postgres":
	// 	plugin, err = installPostgres()
	// case "mongo":
	// 	plugin, err = installMongo()
	// case "redis":
	// 	plugin, err = installRedis()
	default:
		err = fmt.Errorf("%s is not a valid plugin name", pluginName)
	}

	return
}

func installTraefik(tx *gorm.DB) (*models.Plugin, *server_models.AppDetails, error) {
	log.V("plugin_server.installTraefik()")
	if IsInstalled(tx, constants.PluginNameTraefik) {
		return nil, nil, fmt.Errorf("traefik is already installed")
	}

	err := UpdatePluginInstalled(tx, constants.PluginNameTraefik, true)
	if err != nil {
		return nil, nil, err
	}
	traefik := constants.Plugins.Traefik

	err = docker_service.PullImage(traefik.App.Image)
	if err != nil {
		return nil, nil, err
	}
	digest, err := docker_service.GetImageDigest(traefik.App.Image)
	if err != nil {
		return nil, nil, err
	}
	traefik.RunConfig.ImageDigest = digest

	// TODO: Add traefik config for port 4000, pass into app_mapper.ToService

	plugin, err := Get(tx, constants.PluginNameTraefik)
	if err != nil {
		return nil, nil, err
	}
	return plugin, traefik, err
}
