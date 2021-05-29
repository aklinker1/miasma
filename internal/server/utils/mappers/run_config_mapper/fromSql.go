package run_config_mapper

import (
	"encoding/json"

	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/package/models"
)

func FromSQL(sql *entities.SQLRunConfig) *models.RunConfig {
	volumes := []*models.RunConfigVolumesItems0{}
	jsonVolumes := []*entities.SQLRunConfigVolume{}
	err := json.Unmarshal(sql.Volumes, &jsonVolumes)
	if err != nil {
		panic(err)
	}
	for _, volume := range jsonVolumes {
		volumes = append(volumes, &models.RunConfigVolumesItems0{
			Source: volume.Source,
			Target: volume.Target,
		})
	}

	return &models.RunConfig{
		AppID:          sql.AppID,
		Command:        sql.Command,
		ImageDigest:    sql.ImageDigest,
		Networks:       sql.Networks,
		Placement:      sql.Networks,
		PublishedPorts: sql.PublishedPorts,
		TargetPorts:    sql.PublishedPorts,
		Volumes:        volumes,
	}
}
