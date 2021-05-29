package handlers

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/services/run_config_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/aklinker1/miasma/package/models"
	"github.com/go-openapi/runtime/middleware"
)

var CreateApp = operations.CreateAppHandlerFunc(
	func(params operations.CreateAppParams) middleware.Responder {
		log.V("handlers.CreateApp()")
		var err error
		tx, onDefer := database.TX(&err)
		defer onDefer()

		if params.App == nil {
			return operations.NewCreateAppBadRequest().WithPayload("Request body is required")
		}
		inputApp := *params.App
		err = validation.AppName(inputApp.Name)
		if err != nil {
			return operations.NewCreateAppBadRequest().WithPayload(err.Error())
		}
		existingApp, _ := app_service.Get(tx, inputApp.Name)
		if existingApp != nil {
			log.D("Existing app: %v", existingApp)
			return operations.NewCreateAppBadRequest().WithPayload(fmt.Sprintf("%s already exists", inputApp.Name))
		}

		err = docker_service.PullImage(inputApp.Image)
		if err != nil {
			return operations.NewCreateAppBadRequest().WithPayload(err.Error())
		}
		digest, err := docker_service.GetImageDigest(inputApp.Image)
		if err != nil {
			return operations.NewCreateAppDefault(500).WithPayload(err.Error())
		}

		newApp, err := app_service.Create(tx, inputApp)
		if err != nil {
			return operations.NewCreateAppDefault(500).WithPayload(err.Error())
		}

		defaultRunConfig := &models.RunConfig{
			AppID:       newApp.ID,
			ImageDigest: digest,
		}
		err = run_config_service.Create(tx, defaultRunConfig)
		if err != nil {
			return operations.NewCreateAppDefault(500).WithPayload(err.Error())
		}

		return operations.NewCreateAppCreated().WithPayload(newApp)
	},
)
