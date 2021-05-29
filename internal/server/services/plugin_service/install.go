package plugin_service

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
)

func Install(pluginName string) (plugin *models.Plugin, err error) {
	switch pluginName {
	case "traefik":
		plugin, err = installTreafik()
	// case "postgres":
	// 	plugin, err = installPostgres()
	// case "mongo":
	// 	plugin, err = installMongo()
	// case "redis":
	// 	plugin, err = installRedis()
	default:
		err = fmt.Errorf("%s is not a valid plugin name", pluginName)
	}

	return plugin, err
}

func installTreafik() (*models.Plugin, error) {
	panic("TODO: Save that the app is installed")
	// pluginMeta.Traefik = true
	traefik := constants.Plugins.Traefik

	err := docker_service.PullImage(traefik.App.Image)
	if err != nil {
		return nil, err
	}
	digest, err := docker_service.GetImageDigest(traefik.App.Image)
	if err != nil {
		return nil, err
	}
	log.W("TODO: Remove this log %v", digest)

	panic("TODO: Add traefik config for port 4000, pass into app_mapper.ToService")
	// plugins := &server_models.AppPlugins{}

	// env := map[string]string{}
	// serviceSpec, _ := app_mapper.ToService(
	// 	traefik,
	// 	env,
	// 	plugins,
	// 	func(i int) ([]uint32, error) {
	// 		return traefik.RunConfig.PublishedPorts, nil
	// 	},
	// 	digest,
	// )
	// serviceOptions := types.ServiceCreateOptions{}

	// err = docker_service.UpsertNetwork(traefik.App.Name)
	// if err != nil {
	// 	return nil, err
	// }
	// err = docker_service.StartService(*serviceSpec, serviceOptions)
	// if err != nil {
	// 	return nil, err
	// }
	// err = WritePluginMeta(pluginMeta)
	// if err != nil {
	// 	return nil, err
	// }
	// return Get("traefik", pluginMeta)
}
