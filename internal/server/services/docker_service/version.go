package docker_service

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/shared/log"
)

func Version() *string {
	log.V("docker_service.Version()")
	version, err := dockerAPI.ServerVersion(ctx)
	if err != nil {
		log.E("%v", err)
		return nil
	}
	versionString := fmt.Sprintf("%s-%s", version.Version, version.GitCommit)
	return &versionString
}
