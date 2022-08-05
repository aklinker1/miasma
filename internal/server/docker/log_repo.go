package docker

import (
	"context"

	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/io"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type logRepo struct {
	logger server.Logger
	client *client.Client
}

func NewLogRepo(logger server.Logger, client *client.Client) server.LogRepo {
	return &logRepo{
		logger: logger,
		client: client,
	}
}

// GetLogStream implements server.LogRepo
func (r *logRepo) GetLogStream(ctx context.Context, serviceID string) (server.LogStream, error) {
	rd, err := r.client.ServiceLogs(ctx, serviceID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     true,
		Tail:       "50",
	})
	if err != nil {
		return nil, err
	}
	return io.NewLogStream(rd), nil
}
