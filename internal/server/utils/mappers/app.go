package mappers

import (
	dockerSwarmTypes "docker.io/go-docker/api/types/swarm"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/types"
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

func (a *app) ToMeta(app *models.App) *types.AppMetaData {
	return &types.AppMetaData{
		Image:  app.Image,
		Hidden: &app.Hidden,
	}
}

func (a *app) ToService(app *models.App) *dockerSwarmTypes.ServiceSpec {
	return &dockerSwarmTypes.ServiceSpec{
		TaskTemplate: dockerSwarmTypes.TaskSpec{
			Placement: &dockerSwarmTypes.Placement{
				Constraints: []string{},
			},
			ContainerSpec: &dockerSwarmTypes.ContainerSpec{
				Image: *app.Image,
				// Env: ,
			},
		},
		Annotations: dockerSwarmTypes.Annotations{
			Name: *app.Name,
		},
	}
}
