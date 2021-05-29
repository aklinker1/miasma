package run_config_mapper

import (
	"encoding/json"

	"github.com/aklinker1/miasma/internal/server/database/entities"
	"github.com/aklinker1/miasma/package/models"
)

func FromSQL(sql *entities.SQLRunConfig) *models.RunConfig {
	volumes := []*models.RunConfigVolume{}
	jsonVolumes := []*entities.SQLRunConfigVolume{}
	err := json.Unmarshal(sql.Volumes, &jsonVolumes)
	if err != nil {
		panic(err)
	}
	for _, volume := range jsonVolumes {
		volumes = append(volumes, &models.RunConfigVolume{
			Source: volume.Source,
			Target: volume.Target,
		})
	}
	if len(volumes) == 0 {
		volumes = nil
	}

	return &models.RunConfig{
		AppID:          sql.AppID,
		Command:        sql.Command,
		ImageDigest:    sql.ImageDigest,
		Networks:       sql.Networks,
		Placement:      sql.Placement,
		PublishedPorts: sql.PublishedPorts,
		TargetPorts:    sql.TargetPorts,
		Volumes:        volumes,
	}
}
