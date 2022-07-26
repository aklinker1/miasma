package docker

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type pullImageStatus struct {
	Status string `json:"status"`
}

type runtimeImageRepo struct {
	client client.APIClient
	logger server.Logger
}

func NewRuntimeImageRepo(logger server.Logger, client *client.Client) server.RuntimeImageRepo {
	return &runtimeImageRepo{
		client: client,
		logger: logger,
	}
}

// PullLatest implements server.RuntimeImageRepo
func (s *runtimeImageRepo) GetLatestDigest(ctx context.Context, image string) (string, error) {
	s.logger.D("Pulling latest image: %s", image)
	stream, err := s.client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		s.logger.E("Failed to pull image: %v", err)
		return "", err
	}
	defer stream.Close()

	// Read each line separately, they each return JSON: { "status": "..." }
	var digest string
	reader := bufio.NewReader(stream)
	for true {
		data, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.logger.E("Failed to read line: %v", err)
			return "", err
		}

		var status pullImageStatus
		err = json.Unmarshal(data, &status)
		if err != nil {
			return "", err
		}
		s.logger.V(status.Status)
		if strings.HasPrefix(status.Status, "Digest:") {
			digest = strings.TrimSpace(strings.ReplaceAll(status.Status, "Digest: ", ""))
		}
	}

	if digest == "" {
		return "", &server.Error{
			Code:    server.EINTERNAL,
			Message: "Image pull did not report the digest",
			Op:      "docker.runtimeImageRepo.PullImage",
		}
	}
	return digest, nil
}
