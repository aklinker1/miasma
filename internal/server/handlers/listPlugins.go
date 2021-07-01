package handlers

import (
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/plugin_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/go-openapi/runtime/middleware"
)

var ListPlugins = operations.ListPluginsHandlerFunc(
	func(params operations.ListPluginsParams) middleware.Responder {
		log.V("handlers.ListPlugins()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		plugins := plugin_service.List(db)
		return operations.NewListPluginsOK().WithPayload(plugins)
	},
)
