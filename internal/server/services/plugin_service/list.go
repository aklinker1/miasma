package plugin_service

import (
	"sort"

	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/package/models"
)

func List() (plugins []*models.Plugin, err error) {
	pluginMap := shared.StructToMap(constants.Plugins)
	names := []string{}
	for pluginName := range pluginMap {
		names = append(names, pluginName)
	}
	sort.Strings(names)
	panic("TODO: List installed statuses for available plugins and return them")
	// pluginMeta, err := ListInstallState(names)
	// if err != nil {
	// 	return nil, err
	// }
	// for _, pluginName := range names {
	// 	plugin, err := Get(pluginName, pluginMeta)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	plugins = append(plugins, plugin)
	// }

	// return plugins, nil
}
