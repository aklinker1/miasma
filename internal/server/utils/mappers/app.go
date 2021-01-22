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
		Plugins:     []string{},
		Env:         map[string]string{},
	}
}

func (a *app) ToConfig(app *types.AppMetaData) *models.AppConfig {
	return &models.AppConfig{
		TargetPorts: shared.ConvertUInt32ArrayToInt64Array(app.TargetPorts),
		// Networks: ,
		// Plugins: ,
	}
}

func (a *app) ToService(app *models.App, getNextPort func() (uint32, error)) (*dockerSwarmTypes.ServiceSpec, error) {
	// Get ports
	defaultPort := shared.RandUInt32(3000, 4000)
	targetPorts := append([]uint32{defaultPort}, []uint32{}...)
	envPorts := []string{fmt.Sprintf("PORT=%d", defaultPort)}
	portConfigs := []dockerSwarmTypes.PortConfig{}
	for index, targetPort := range targetPorts {
		nextPublishedPort, err := getNextPort()
		if err != nil {
			return nil, err
		}
		envPort := fmt.Sprintf("PORT_%d=%d", index+1, targetPort)
		envPorts = append(envPorts, envPort)
		portConfigs = append(portConfigs, dockerSwarmTypes.PortConfig{
			PublishedPort: nextPublishedPort,
			TargetPort:    targetPort,
		})
		if index == 0 {
			log.V("Ports Env: %d:%d (%s, %s)", nextPublishedPort, targetPort, envPorts[0], envPort)
		} else {
			log.V("Ports Env: %d:%d (%s)", nextPublishedPort, targetPort, envPort)
		}
	}

	// Setup env variables
	env := append(envPorts, []string{}...)

	return &dockerSwarmTypes.ServiceSpec{
		Annotations: dockerSwarmTypes.Annotations{
			Name: *app.Name,
		},
		TaskTemplate: dockerSwarmTypes.TaskSpec{
			Placement: &dockerSwarmTypes.Placement{
				Constraints: []string{},
			},
			ContainerSpec: &dockerSwarmTypes.ContainerSpec{
				Image: *app.Image,
				Env:   env,
			},
			// Networks: ,
			RestartPolicy: &dockerSwarmTypes.RestartPolicy{
				Condition: dockerSwarmTypes.RestartPolicyConditionOnFailure,
			},
		},
		EndpointSpec: &dockerSwarmTypes.EndpointSpec{
			Ports: portConfigs,
		},
	}, nil
}
