package run_config_mapper

import (
	"encoding/json"

	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/package/models"
)

func ToSQL(model *models.RunConfig) *entities.SQLRunConfig {
	volumes := []*entities.SQLRunConfigVolume{}
	for _, volume := range model.Volumes {
		volumes = append(volumes, &entities.SQLRunConfigVolume{
			Source: volume.Source,
			Target: volume.Target,
		})
	}
	volumeBytes, err := json.Marshal(volumes)
	if err != nil {
		panic(err)
	}

	return &entities.SQLRunConfig{
		AppID:          model.AppID,
		Command:        model.Command,
		ImageDigest:    model.ImageDigest,
		Networks:       model.Networks,
		Placement:      model.Networks,
		PublishedPorts: model.PublishedPorts,
		TargetPorts:    model.PublishedPorts,
		Volumes:        volumeBytes,
	}
}
