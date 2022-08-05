package docker

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/io"
	"github.com/aklinker1/miasma/internal/utils"
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
func (r *logRepo) GetLogStream(ctx context.Context, filter server.LogsFilter) (server.LogStream, error) {
	showStdout := !utils.ValueOr(filter.ExcludeStdout, false)
	showStderr := !utils.ValueOr(filter.ExcludeStderr, false)
	if !showStdout && !showStderr {
		return nil, fmt.Errorf("excludeStdout and excludeStderr cannot both be true")
	}
	options := types.ContainerLogsOptions{
		Timestamps: true,
		ShowStdout: showStdout,
		ShowStderr: showStderr,
	}
	if filter.After != nil {
		options.Since = fmt.Sprint(filter.After.Unix())
	}
	if filter.Before != nil {
		options.Until = fmt.Sprint(filter.Before.Unix())
	}
	if filter.Follow != nil {
		options.Follow = true
	}
	if filter.Tail != nil {
		options.Tail = *filter.Tail
	}

	rd, err := r.client.ServiceLogs(ctx, filter.ServiceID, options)
	if err != nil {
		return nil, err
	}
	return io.NewLogStream(rd), nil
}
