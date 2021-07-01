package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/aklinker1/miasma/package/models"
	"github.com/go-openapi/runtime/middleware"
)

var UpdateAppTraefikConfig = operations.UpdateAppTraefikConfigHandlerFunc(
	func(params operations.UpdateAppTraefikConfigParams) middleware.Responder {
		log.V("handlers.UpdateAppTraefikConfig()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		err = validation.TraefikPluginConfig(params.NewTraefikConfig)
		if err != nil {
			return operations.NewUpdateAppTraefikConfigBadRequest().WithPayload(err.Error())
		}
		if !plugin_service.IsInstalled(tx, constants.PluginNameTraefik) {
			return operations.NewUpdateAppTraefikConfigBadRequest().WithPayload("Traefik is not installed")
		}

		config := &models.TraefikPluginConfig{
			AppID:       params.AppID,
			Host:        params.NewTraefikConfig.Host,
			Path:        params.NewTraefikConfig.Path,
			TraefikRule: params.NewTraefikConfig.TraefikRule,
		}
		err = plugin_service.UpsertTraefikConfig(tx, config)
		if err != nil {
			return operations.NewUpdateAppTraefikConfigDefault(500).WithPayload(err.Error())
		}

		return operations.NewUpdateAppTraefikConfigOK().WithPayload(config)
	},
)
