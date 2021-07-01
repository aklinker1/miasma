package docker_service

import "github.com/aklinker1/miasma/internal/shared/log"

func DeleteNetwork(appName string) error {
	log.V("docker_service.deleteNetwork")
	return dockerAPI.NetworkRemove(ctx, appName)
}
