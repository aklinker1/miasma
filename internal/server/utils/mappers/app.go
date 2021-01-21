package mappers

import (
	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/types"
)

type app struct{}

var App = &app{}

func (a *app) FromMeta(appName string, meta *types.AppMetaData) *models.App {
	return &models.App{
		Name:   &appName,
		Image:  meta.Image,
		Hidden: meta.Hidden != nil && *meta.Hidden,
	}
}

func (a *app) ToMeta(app *models.App) *types.AppMetaData {
	return &types.AppMetaData{
		Image:  app.Image,
		Hidden: &app.Hidden,
	}
}
