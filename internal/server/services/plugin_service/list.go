package plugin_service

import (
	"sort"

	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/package/models"
	"gorm.io/gorm"
)

func List(tx *gorm.DB) []*models.Plugin {
	pluginMap := shared.StructToMap(constants.Plugins)
	names := []string{}
	for pluginName := range pluginMap {
		names = append(names, pluginName)
	}
	sort.Strings(names)

	plugins := []*models.Plugin{}
	for _, name := range names {
		plugins = append(plugins, &models.Plugin{
			Name:      name,
			Installed: IsInstalled(tx, name),
		})
	}
	return plugins
}
