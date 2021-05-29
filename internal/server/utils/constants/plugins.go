package constants

import (
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/package/models"
)

type predefinedPlugins struct {
	Traefik *server_models.AppDetails `json:"traefik"`
}

const (
	traefikID = "ff828d53-f29a-4c42-a426-e465f9005be6"
)

var Plugins = predefinedPlugins{
	// https://doc.traefik.io/traefik/getting-started/quick-start/
	Traefik: &server_models.AppDetails{
		App: &models.App{
			Hidden: false,
			ID:     traefikID,
			Image:  "traefik:v2.4",
			Name:   "plugin-traefik",
		},
		RunConfig: &models.RunConfig{
			AppID:          traefikID,
			TargetPorts:    []uint32{80, 8080},
			PublishedPorts: []uint32{80, 4000},
			Placement: []string{
				"node.labels.traefik==true",
				"node.role==manager",
			},
			Volumes: []*models.RunConfigVolumesItems0{{
				Source: "/var/run/docker.sock",
				Target: "/var/run/docker.sock",
			}},
			Command: []string{"traefik", "--api.insecure=true", "--providers.docker", "--providers.docker.swarmmode"},
		},
	},
}
