package docker_service

import (
	"context"

	"docker.io/go-docker"
	"github.com/aklinker1/miasma/internal/shared/log"
)

var dockerAPI *docker.Client
var ctx = context.Background()

func init() {
	log.V("docker_service.init")
	var err error
	dockerAPI, err = docker.NewEnvClient()
	if err != nil {
		panic("Could not connect to host's docker service")
	}
}
