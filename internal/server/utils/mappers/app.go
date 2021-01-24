package mappers

import (
	"fmt"
	"strings"

	"docker.io/go-docker/api/types/mount"
	dockerSwarmTypes "docker.io/go-docker/api/types/swarm"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/internal/shared/log"
)

type app struct{}

var App = &app{}

func (a *app) FromMeta(appName string, meta *types.AppMetaData, isRunning bool) *models.App {
	return &models.App{
		Name:    &appName,
		Image:   meta.Image,
		Hidden:  meta.Hidden != nil && *meta.Hidden,
		Running: &isRunning,
	}
}

func (a *app) ToMeta(app *models.AppInput) *types.AppMetaData {
	return &types.AppMetaData{
		Image:  app.Image,
		Hidden: &app.Hidden,
	}
}

func (a *app) ToConfig(app *types.AppMetaData) *models.AppConfig {
	var route *models.AppConfigRoute
	log.V("App route: %v", app.Route)
	if app.Route != nil {
		if app.Route.TraefikRule != "" {
			route = &models.AppConfigRoute{
				TraefikRule: &app.Route.TraefikRule,
			}
		} else {
			route = &models.AppConfigRoute{
				Host: &app.Route.Host,
				Path: &app.Route.Path,
			}
		}
	}
	return &models.AppConfig{
		TargetPorts:    shared.ConvertUInt32ArrayToInt64Array(app.TargetPorts),
		PublishedPorts: shared.ConvertUInt32ArrayToInt64Array(app.PublishedPorts),
		Placement:      app.Placement,
		Networks:       app.Networks,
		Route:          route,
	}
}

func (a *app) ToService(app *types.AppMetaData, plugins *types.PluginMetaData, getNextPorts func(int) ([]uint32, error)) (*dockerSwarmTypes.ServiceSpec, error) {
	// Setup ports
	var targetPorts = []uint32{}
	if len(app.TargetPorts) > 0 {
		targetPorts = append(targetPorts, app.TargetPorts...)
	} else {
		// Internal to container. This does not have the same limitations (3001-3999) at the
		// published ports. The range being close is a coincidence
		targetPorts = append(targetPorts, shared.RandUInt32(3000, 4000))
	}
	envPorts := []string{fmt.Sprintf("PORT=%d", targetPorts[0])}
	publishedPorts, err := getNextPorts(len(targetPorts))
	if err != nil {
		return nil, err
	}
	portConfigs := []dockerSwarmTypes.PortConfig{}
	for index, targetPort := range targetPorts {
		publishedPort := publishedPorts[index]
		envPort := fmt.Sprintf("PORT_%d=%d", index+1, targetPort)
		envPorts = append(envPorts, envPort)
		portConfigs = append(portConfigs, dockerSwarmTypes.PortConfig{
			PublishedPort: publishedPort,
			TargetPort:    targetPort,
		})
		if index == 0 {
			log.V("Ports Env: %d:%d (%s, %s)", publishedPort, targetPort, envPorts[0], envPort)
		} else {
			log.V("Ports Env: %d:%d (%s)", publishedPort, targetPort, envPort)
		}
	}

	// Setup networks
	networks := []dockerSwarmTypes.NetworkAttachmentConfig{
		{
			Target: app.Name,
		},
	}
	for _, network := range app.Networks {
		log.V("Connecting to another app: %s", network)
		networks = append(networks, dockerSwarmTypes.NetworkAttachmentConfig{
			Target: network,
		})
	}
	if plugins.Traefik && app.Route != nil {
		log.V("Adding traefik network")
		networks = append(networks, dockerSwarmTypes.NetworkAttachmentConfig{
			Target: constants.Plugins.Traefik.Name,
		})
	}
	// TODO: More plugins
	// if plugins.Postgres {
	// 	log.V("Adding postgres network")
	// 	networks = append(networks, dockerSwarmTypes.NetworkAttachmentConfig{
	// 		Target: constants.Plugins.Postgres.Name,
	// 	})
	// }
	// if plugins.Mongo {
	// 	log.V("Adding mongo network")
	// 	networks = append(networks, dockerSwarmTypes.NetworkAttachmentConfig{
	// 		Target: constants.Plugins.Mongo.Name,
	// 	})
	// }

	// Setup env variables
	env := append(envPorts, []string{}...)

	// Setup vVolumes
	volumes := []mount.Mount{}
	for _, volume := range app.Volumes {
		split := strings.Split(volume, ":")
		volumes = append(volumes, mount.Mount{
			Source: split[0],
			Target: split[1],
			Type:   mount.TypeBind,
		})
	}

	// Setup labels
	labels := map[string]string{
		"miasma": "true",
	}
	if app.Route != nil {
		enabled := "traefik.enable"
		networkLabel := "traefik.docker.network"
		targetPort := fmt.Sprintf("traefik.http.services.%s-service.loadbalancer.server.port", app.Name)
		rulesLabel := fmt.Sprintf("traefik.http.routers.%s.rule", app.Name)

		labels[enabled] = "true"
		labels[networkLabel] = constants.Plugins.Traefik.Name
		labels[targetPort] = fmt.Sprint(targetPorts[0])
		if app.Route.TraefikRule != "" {
			labels[rulesLabel] = app.Route.TraefikRule
		} else if app.Route.Path == "" {
			labels[rulesLabel] = fmt.Sprintf("Host(`%s`)", app.Route.Host)
		} else {
			labels[rulesLabel] = fmt.Sprintf("(Host(`%s`) && PathPrefix(`%s`))", app.Route.Host, app.Route.Path)
		}
	}

	return &dockerSwarmTypes.ServiceSpec{
		Annotations: dockerSwarmTypes.Annotations{
			Name:   app.Name,
			Labels: labels,
		},
		TaskTemplate: dockerSwarmTypes.TaskSpec{
			Placement: &dockerSwarmTypes.Placement{
				Constraints: app.Placement,
			},
			ContainerSpec: &dockerSwarmTypes.ContainerSpec{
				Image:   *app.Image,
				Env:     env,
				Command: app.Command,
				Mounts:  volumes,
			},
			Networks: networks,
		},
		EndpointSpec: &dockerSwarmTypes.EndpointSpec{
			Ports: portConfigs,
		},
	}, nil
}
