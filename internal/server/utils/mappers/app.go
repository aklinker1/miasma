package mappers

import (
	"fmt"

	dockerSwarmTypes "docker.io/go-docker/api/types/swarm"

	"github.com/aklinker1/miasma/internal/server/gen/models"
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
		Image:       app.Image,
		Hidden:      &app.Hidden,
		TargetPorts: []uint32{},
		Networks:    []string{},
		Env:         map[string]string{},
	}
}

func (a *app) ToConfig(app *types.AppMetaData) *models.AppConfig {
	return &models.AppConfig{
		TargetPorts: shared.ConvertUInt32ArrayToInt64Array(app.TargetPorts),
		Placement:   app.Placement,
		Networks:    app.Networks,
	}
}

func (a *app) ToService(app *types.AppMetaData, getNextPorts func(int) ([]uint32, error)) (*dockerSwarmTypes.ServiceSpec, error) {
	// Get ports
	var targetPorts = []uint32{}
	if len(app.TargetPorts) > 0 {
		targetPorts = append(targetPorts, app.TargetPorts...)
	} else {
		defaultPort := shared.RandUInt32(3000, 4000)
		targetPorts = append(targetPorts, defaultPort)
	}
	envPorts := []string{fmt.Sprintf("PORT=%d", targetPorts[0])}
	publishedPorts, err := getNextPorts(len(targetPorts))
	if err != nil {
		return nil, err
	}
	portConfigs := []dockerSwarmTypes.PortConfig{}
	for index, targetPort := range targetPorts {
		publishedPort := publishedPorts[index]
		if err != nil {
			return nil, err
		}
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

	// Setup Networks
	networks := []dockerSwarmTypes.NetworkAttachmentConfig{
		{
			Target: app.Name,
		},
	}
	for _, network := range app.Networks {
		log.V("Additional network: %s", network)
		networks = append(networks, dockerSwarmTypes.NetworkAttachmentConfig{
			Target: network,
		})
	}

	// Setup env variables
	env := append(envPorts, []string{}...)

	return &dockerSwarmTypes.ServiceSpec{
		Annotations: dockerSwarmTypes.Annotations{
			Name: app.Name,
		},
		TaskTemplate: dockerSwarmTypes.TaskSpec{
			Placement: &dockerSwarmTypes.Placement{
				Constraints: app.Placement,
			},
			ContainerSpec: &dockerSwarmTypes.ContainerSpec{
				Image: *app.Image,
				Env:   env,
			},
			Networks: networks,
		},
		EndpointSpec: &dockerSwarmTypes.EndpointSpec{
			Ports: portConfigs,
		},
	}, nil
}
