package constants

import (
	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared"
)

type predefinedPlugins struct {
	Traefik types.AppMetaData `json:"traefik"`
}

var Plugins = predefinedPlugins{
	// https://doc.traefik.io/traefik/getting-started/quick-start/
	Traefik: types.AppMetaData{
		Name:           "plugin-traefik",
		Image:          shared.StringPtr("traefik:v2.4"),
		TargetPorts:    []uint32{80, 8080},
		PublishedPorts: []uint32{80, 4000},
		Placement:      []string{"node.labels.traefik==true"},
		Volumes: []*models.AppConfigVolumesItems0{{
			Source: "/var/run/docker.sock",
			Target: "/var/run/docker.sock",
		}},
		Command: []string{"traefik", "--api.insecure=true", "--providers.docker", "--providers.docker.swarmmode"},
	},
}
