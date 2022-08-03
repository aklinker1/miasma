package docker

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

type runtimeNodeRepo struct {
	client client.APIClient
	logger server.Logger
}

func NewRuntimeNodeRepo(logger server.Logger, client *client.Client) server.RuntimeNodeRepo {
	return &runtimeNodeRepo{
		client: client,
		logger: logger,
	}
}

// ListNodes implements server.RuntimeNodeRepo
func (s *runtimeNodeRepo) GetAll(ctx context.Context, filter server.RuntimeNodesFilter) ([]internal.Node, error) {
	nodes, err := s.client.NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		return nil, err
	}

	return lo.Map(nodes, func(n swarm.Node, _ int) internal.Node {
		var statusMessage *string
		if n.Status.Message != "" {
			statusMessage = lo.ToPtr(n.Status.Message)
		}
		return internal.Node{
			ID:            n.ID,
			Os:            n.Description.Platform.OS,
			Architecture:  n.Description.Platform.Architecture,
			Status:        string(n.Status.State),
			StatusMessage: statusMessage,
			Hostname:      n.Description.Hostname,
			IP:            n.Status.Addr,
			Labels:        utils.ToAnyMap(n.Spec.Labels),
		}
	}), nil
}

// GetOne implements server.RuntimeNodeRepo
func (s *runtimeNodeRepo) GetOne(ctx context.Context, filter server.RuntimeNodesFilter) (internal.Node, error) {
	nodes, err := s.GetAll(ctx, filter)
	if err != nil {
		return zero.Node, err
	}
	if len(nodes) == 0 {
		return zero.Node, &server.Error{
			Code:    server.ENOTFOUND,
			Message: fmt.Sprintf("Node not found for %+v", filter),
		}
	}
	return nodes[0], nil
}
