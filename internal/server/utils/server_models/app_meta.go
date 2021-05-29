package server_models

import "github.com/aklinker1/miasma/package/models"

type AppDetails struct {
	App       *models.App
	RunConfig *models.RunConfig
}
