package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var RemoveAppTraefikConfig = operations.RemoveAppTraefikConfigHandlerFunc(
	func(params operations.RemoveAppTraefikConfigParams) middleware.Responder {
		log.V("handlers.RemoveAppTraefikConfig()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		if !plugin_service.IsInstalled(tx, constants.PluginNameTraefik) {
			return operations.NewRemoveAppTraefikConfigBadRequest().WithPayload("Traefik is not installed")
		}
		appID := params.AppID.String()

		plugins, err := plugin_service.GetAppPlugins(tx, appID)
		if plugins.Traefik == nil {
			return operations.NewRemoveAppTraefikConfigBadRequest().
				WithPayload(fmt.Sprintf("No traefik config to delete for '%v'", appID))
		}
		if err != nil {
			return operations.NewRemoveAppTraefikConfigDefault(500).WithPayload(err.Error())
		}

		err = plugin_service.DeleteTraefikConfig(tx, appID)
		if err != nil {
			return operations.NewRemoveAppTraefikConfigDefault(500).WithPayload(err.Error())
		}

		return operations.NewRemoveAppTraefikConfigOK().WithPayload(plugins.Traefik)
	},
)
