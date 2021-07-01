package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var GetAppTraefikConfig = operations.GetAppTraefikConfigHandlerFunc(
	func(params operations.GetAppTraefikConfigParams) middleware.Responder {
		log.V("handlers.GetAppTraefikConfig()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		if !plugin_service.IsInstalled(db, constants.PluginNameTraefik) {
			return operations.NewGetAppTraefikConfigBadRequest().
				WithPayload("Traefik is not installed")
		}

		appID := params.AppID.String()
		plugins, err := plugin_service.GetAppPlugins(db, appID)
		if err != nil {
			return operations.NewGetAppTraefikConfigDefault(500).WithPayload(err.Error())
		}
		traefikConfig := plugins.Traefik
		if traefikConfig == nil {
			return operations.NewGetAppTraefikConfigNotFound()
		}

		return operations.NewGetAppTraefikConfigOK().WithPayload(traefikConfig)
	},
)
