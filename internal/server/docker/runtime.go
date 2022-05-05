package docker

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/client"
)

type RuntimeService struct {
	client client.APIClient
	logger server.Logger
}

func NewRuntimeService(logger server.Logger) (server.RuntimeService, error) {
	c, err := client.NewClientWithOpts(client.FromEnv)
	return &RuntimeService{
		client: c,
		logger: logger,
	}, err
}

// PullLatest implements server.RuntimeService
func (*RuntimeService) PullLatest(ctx context.Context, image string) (string, error) {
	return "", server.NewNotImplementedError("docker.RuntimeService.PullLatest")
}

// Restart implements server.RuntimeService
func (*RuntimeService) Restart(ctx context.Context, app internal.App) error {
	return server.NewNotImplementedError("docker.RuntimeService.Restart")
}

// Start implements server.RuntimeService
func (*RuntimeService) Start(ctx context.Context, app internal.App) error {
	return server.NewNotImplementedError("docker.RuntimeService.Start")
}

// Stop implements server.RuntimeService
func (*RuntimeService) Stop(ctx context.Context, app internal.App) error {
	return server.NewNotImplementedError("docker.RuntimeService.Stop")
}

// SwarmInfo implements server.RuntimeService
func (s *RuntimeService) SwarmInfo(ctx context.Context) (*internal.SwarmInfo, error) {
	info, err := s.client.Info(ctx)
	if err != nil {
		return nil, &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to run 'docker info'",
			Op:      "docker.RuntimeService.SwarmInfo()",
			Err:     err,
		}
	}
	swarm, err := s.client.SwarmInspect(ctx)
	if isSwarmNotInitializedError(err) {
		return nil, nil
	} else if err != nil {
		return nil, &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to run 'docker swarm inspect'",
			Op:      "docker.RuntimeService.SwarmInfo()",
			Err:     err,
		}
	}
	return &internal.SwarmInfo{
		ID:          swarm.ID,
		JoinCommand: fmt.Sprintf("docker swarm join --token %s %s:2377", swarm.JoinTokens.Worker, info.Swarm.NodeAddr),
		// JoinCommand: fmt.Sprintf("docker swarm join --token %s <main-node-ip:port>", swarm.JoinTokens.Worker),
		CreatedAt: swarm.CreatedAt,
		UpdatedAt: swarm.UpdatedAt,
	}, nil
}

// Version implements server.RuntimeService
func (s *RuntimeService) Version(ctx context.Context) (string, error) {
	info, err := s.client.Info(ctx)
	return info.ServerVersion, err
}
