package mappers

import (
	"github.com/aklinker1/miasma/internal/gen/models"
	"github.com/aklinker1/miasma/internal/utils/types"
)

type app struct{}

var App = &app{}

func (a *app) FromMeta(appName string, meta *types.AppMetaData) *models.App {
	return &models.App{
		Name:  &appName,
		Image: meta.Image,
	}
}
