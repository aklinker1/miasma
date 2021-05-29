package run_config_service

import (
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func GetNextPorts(details *server_models.AppDetails) func(int) ([]uint32, error) {
	return func(portCountNeeded int) ([]uint32, error) {
		// We have published ports and they are enough to meet the request, return them
		if len(details.RunConfig.PublishedPorts) > 0 && len(details.RunConfig.PublishedPorts) >= portCountNeeded {
			return details.RunConfig.PublishedPorts, nil
		}

		// otherwise start the final list of ports with the published ports, adding any additional
		// ports from the existing service, and finally get the next open ones if more are necessary
		newPublishedPorts := details.RunConfig.PublishedPorts

		existingService, _ := docker_service.GetRunningService(details.App.Name)
		existingPublishedPorts := []uint32{}
		if existingService != nil {
			for _, port := range existingService.Spec.EndpointSpec.Ports {
				log.V("Old port: -p %d:%d", port.PublishedPort, port.TargetPort)
				existingPublishedPorts = append(existingPublishedPorts, port.PublishedPort)
			}
		}
		// Adding only extra, already published ports so those stay the same
		if len(existingPublishedPorts) > len(newPublishedPorts) {
			newPublishedPorts = append(newPublishedPorts, existingPublishedPorts[len(newPublishedPorts):]...)
		}

		// if more are still needed, find the next set of open ports
		if len(newPublishedPorts) < portCountNeeded {
			numberOfAdditionalPortsNeeded := portCountNeeded - len(newPublishedPorts)
			nextOpenPorts, err := docker_service.GetNextFreePorts(numberOfAdditionalPortsNeeded)
			if err != nil {
				return nil, err
			}
			newPublishedPorts = append(newPublishedPorts, nextOpenPorts...)
		}
		log.V("Ports after update: %v", newPublishedPorts)
		return newPublishedPorts, nil
	}
}
