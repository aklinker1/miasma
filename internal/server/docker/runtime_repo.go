package docker

import (
	"context"
	"strings"

	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type runtimeRepo struct {
	client client.APIClient
	logger server.Logger
}

func NewRuntimeRepo(logger server.Logger, client *client.Client) server.RuntimeRepo {
	return &runtimeRepo{
		client: client,
		logger: logger,
	}
}

// Version implements server.RuntimeRepo
func (s *runtimeRepo) Info(ctx context.Context) (types.Info, error) {
	s.logger.D("Getting docker info")
	return s.client.Info(ctx)
}

// ClusterInfo implements server.RuntimeRepo
func (s *runtimeRepo) ClusterInfo(ctx context.Context) (*swarm.Swarm, error) {
	s.logger.D("Getting cluster info")

	swarm, err := s.client.SwarmInspect(ctx)
	if isSwarmNotInitializedError(err) {
		return nil, nil
	} else if err != nil {
		return nil, &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to run 'docker swarm inspect'",
			Op:      "docker.runtimeRepo.ClusterInfo()",
			Err:     err,
		}
	}
	return &swarm, nil
}

func isSwarmNotInitializedError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "This node is not a swarm manager")
}
