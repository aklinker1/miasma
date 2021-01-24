package services

import (
	"fmt"
	"io/ioutil"
	"sort"

	dockerTypes "docker.io/go-docker/api/types"
	"gopkg.in/yaml.v2"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/server/utils/mappers"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/internal/shared/log"
)

type pluginService struct{}

var Plugin = &pluginService{}

var pluginMetaPath = "/data/miasma/plugins.yml"

func (service *pluginService) GetPluginMeta() (*types.PluginMetaData, error) {
	metaFile, err := ioutil.ReadFile(pluginMetaPath)
	if err != nil {
		return nil, fmt.Errorf("Could not find plugins.yml, was it deleted?", pluginMetaPath)
	}

	var metaYml = &types.PluginMetaData{}
	if err := yaml.Unmarshal(metaFile, metaYml); err != nil {
		return nil, err
	}
	return metaYml, nil
}

func (service *pluginService) WritePluginMeta(pluginMeta *types.PluginMetaData) error {
	data, err := yaml.Marshal(pluginMeta)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pluginMetaPath, data, 0755)
}

func (service *pluginService) Get(pluginName string, meta *types.PluginMetaData) (*models.Plugin, error) {
	installed := false
	switch pluginName {
	case "traefik":
		installed = meta.Traefik
	case "postgres":
		installed = meta.Postgres
	case "mongo":
		installed = meta.Mongo
	}

	var installCommand *string
	var uninstallCommand *string
	if installed {
		uninstallCommand = shared.StringPtr(fmt.Sprintf("miasma plugin:uninstall %s", pluginName))
	} else {
		installCommand = shared.StringPtr(fmt.Sprintf("miasma plugin:install %s", pluginName))
	}

	return &models.Plugin{
		Name:             &pluginName,
		Installed:        &installed,
		InstallCommand:   installCommand,
		UninstallCommand: uninstallCommand,
	}, nil
}

func (service *pluginService) ListAll() (plugins []*models.Plugin, err error) {
	pluginMap := shared.StructToMap(constants.Plugins)
	names := []string{}
	for pluginName := range pluginMap {
		names = append(names, pluginName)
	}
	sort.Strings(names)
	pluginMeta, err := service.GetPluginMeta()
	if err != nil {
		return nil, err
	}
	for _, pluginName := range names {
		plugin, err := service.Get(pluginName, pluginMeta)
		if err != nil {
			return nil, err
		}
		plugins = append(plugins, plugin)
	}

	return plugins, nil
}

func (service *pluginService) Install(pluginName string) (plugin *models.Plugin, err error) {
	pluginMeta, err := service.GetPluginMeta()
	if err != nil {
		return nil, err
	}

	switch pluginName {
	case "traefik":
		plugin, err = service.installTreafik(pluginMeta)
	// case "postgres":
	// 	plugin, err = service.installPostgres()
	// case "mongo":
	// 	plugin, err = service.installMongo()
	// case "redis":
	// 	plugin, err = service.installRedis()
	default:
		err = fmt.Errorf("%s is not a valid plugin name", pluginName)
	}

	return plugin, err
}

func (service *pluginService) installTreafik(pluginMeta *types.PluginMetaData) (*models.Plugin, error) {
	pluginMeta.Traefik = true
	traefik := constants.Plugins.Traefik
	serviceSpec, _ := mappers.App.ToService(&traefik, pluginMeta, func(i int) ([]uint32, error) {
		return traefik.PublishedPorts, nil
	})
	serviceOptions := dockerTypes.ServiceCreateOptions{}

	err := Docker.CreateNetworkIfNotAvailable(traefik.Name)
	if err != nil {
		return nil, err
	}
	err = Docker.StartService(*serviceSpec, serviceOptions)
	if err != nil {
		return nil, err
	}
	err = service.WritePluginMeta(pluginMeta)
	if err != nil {
		return nil, err
	}
	return service.Get("traefik", pluginMeta)
}

func (service *pluginService) Uninstall(pluginName string) (plugin *models.Plugin, err error) {
	pluginMeta, err := service.GetPluginMeta()
	if err != nil {
		return nil, err
	}

	switch pluginName {
	case "traefik":
		plugin, err = service.uninstallTreafik(pluginMeta)
	// case "postgres":
	// 	plugin, err = service.installPostgres(pluginMeta)
	// case "mongo":
	// 	plugin, err = service.installMongo(pluginMeta)
	// case "redis":
	// 	plugin, err = service.installRedis(pluginMeta)
	default:
		err = fmt.Errorf("%s is not a valid plugin name", pluginName)
	}

	return plugin, err
}

func (service *pluginService) uninstallTreafik(pluginMeta *types.PluginMetaData) (*models.Plugin, error) {
	pluginMeta.Traefik = false
	traefik := constants.Plugins.Traefik

	err := Docker.StopService(traefik.Name)
	if err != nil {
		return nil, err
	}
	err = Docker.DestroyNetwork(traefik.Name)
	if err != nil {
		log.W("Failed to destroy network: %v", err)
	}
	_ = service.WritePluginMeta(pluginMeta)
	if err != nil {
		log.W("Failed to mark Traefik as uninstalled: %v", err)
	}

	return service.Get(traefik.Name, pluginMeta)
}
