package docker

import (
	"context"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/client"
)

type RuntimeService struct {
	client client.APIClient
}

func NewRuntimeService() (server.RuntimeService, error) {
	c, err := client.NewClientWithOpts(client.FromEnv)
	return &RuntimeService{
		client: c,
	}, err
}

// PullLatest implements server.RuntimeService
func (*RuntimeService) PullLatest(ctx context.Context, image string) error {
	panic("unimplemented")
}

// Restart implements server.RuntimeService
func (*RuntimeService) Restart(ctx context.Context, app internal.App) error {
	panic("unimplemented")
}

// Start implements server.RuntimeService
func (*RuntimeService) Start(ctx context.Context, app internal.App) error {
	panic("unimplemented")
}

// Stop implements server.RuntimeService
func (*RuntimeService) Stop(ctx context.Context, app internal.App) error {
	panic("unimplemented")
}

// SwarmInfo implements server.RuntimeService
func (*RuntimeService) SwarmInfo(ctx context.Context) error {
	panic("unimplemented")
}
