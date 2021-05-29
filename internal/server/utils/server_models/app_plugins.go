package server_models

import "github.com/aklinker1/miasma/package/models"

type AppPlugins struct {
	Traefik *models.TraefikPluginConfig `json:"traefik"`
}
