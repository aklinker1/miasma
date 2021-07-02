package handlers

import (
	"fmt"

	"docker.io/go-docker/api/types/swarm"
	"github.com/aklinker1/miasma/internal/server/database"
	"github.com/aklinker1/miasma/internal/server/gen/restapi/operations"
	"github.com/aklinker1/miasma/internal/server/services/app_service"
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/shared/log"
	"github.com/aklinker1/miasma/package/models"
	"github.com/go-openapi/runtime/middleware"
)

var ListApps = operations.ListAppsHandlerFunc(
	func(params operations.ListAppsParams) middleware.Responder {
		log.V("handlers.ListApps()")
		db, onDefer := database.ReadOnly()
		defer onDefer()

		showHidden := params.Hidden != nil && *params.Hidden
		apps, err := app_service.List(db, showHidden)
		if err != nil {
			return operations.NewListAppsDefault(500).WithPayload(err.Error())
		}

		services, err := docker_service.RunningServices()
		fmt.Println()
		if err != nil {
			return operations.NewListAppsDefault(500).WithPayload(err.Error())
		}
		serviceMap := map[string]swarm.Service{}
		for _, service := range services {
			serviceMap[service.Spec.Name] = service
		}

		appsWithStatus := make([]*models.AppWithStatus, len(apps))
		for appIndex, app := range apps {
			var ports []string
			var status string
			var instances string

			if service, ok := serviceMap[app.Name]; ok {
				log.V("%s up", app.Name)
				log.V("ports: %v", service.Endpoint.Ports)
				ports = make([]string, len(service.Endpoint.Ports))
				for portIndex, port := range service.Endpoint.Ports {
					ports[portIndex] = fmt.Sprintf(":%d", port.PublishedPort)
				}
				status = "up"
				instances = fmt.Sprintf("%d", *service.Spec.Mode.Replicated.Replicas)
			} else {
				log.V("%s down", app.Name)
				status = "down"
			}
			log.V("ports: %v", ports)
			appsWithStatus[appIndex] = &models.AppWithStatus{
				Name:      app.Name,
				Group:     app.Group,
				Instances: instances,
				Ports:     ports,
				Status:    status,
			}
		}

		return operations.NewListAppsOK().WithPayload(appsWithStatus)
	},
)
