package docker_service

import "github.com/aklinker1/miasma/internal/shared/log"

func StopService(serviceName string) error {
	log.V("docker_service.StopService(%v)", serviceName)
	runningService, err := GetRunningService(serviceName)
	if err != nil {
		return err
	}
	return dockerAPI.ServiceRemove(ctx, runningService.ID)
}
