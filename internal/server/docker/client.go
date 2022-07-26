package docker

import (
	"github.com/docker/docker/client"
)

func NewDefaultClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv)
}
