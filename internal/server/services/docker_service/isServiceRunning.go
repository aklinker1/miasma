package docker_service

import "github.com/aklinker1/miasma/internal/shared/log"

func IsServiceRunning(appName string) bool {
	log.V("docker_service.IsServiceRunning(%v)", appName)
	runningService, _ := GetRunningService(appName)
	return runningService != nil
}
