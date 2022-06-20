package docker

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

var (
	EmptyService               = swarm.Service{}
	EmptyRuntimeServiceDetails = server.RuntimeAppInfo{}
)

var (
	miasmaIdLabel           = "miasma-id"
	miasmaServiceNamePrefix = "miasma-"
)

type pullImageStatus struct {
	Status string `json:"status"`
}

type RuntimeService struct {
	client client.APIClient
	logger server.Logger
}

func NewRuntimeService(logger server.Logger) (server.RuntimeService, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	return &RuntimeService{
		client: client,
		logger: logger,
	}, err
}

// PullLatest implements server.RuntimeService
func (s *RuntimeService) PullLatest(ctx context.Context, image string) (string, error) {
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
			Op:      "docker.RuntimeService.PullImage",
		}
	}
	return digest, nil
}

// Restart implements server.RuntimeService
func (s *RuntimeService) Restart(ctx context.Context, app internal.App) error {
	s.logger.D("Restarting app: %s", app.Name)
	err := s.Stop(ctx, app)
	if err != nil {
		return err
	}
	return s.Start(ctx, app)
}

// Start implements server.RuntimeService
func (s *RuntimeService) Start(ctx context.Context, app internal.App) error {
	s.logger.D("Starting app: %s", app.Name)
	existingService, err := s.getExistingService(ctx, app, false)
	if err != nil {
		return err
	}

	// Ensure the app's network exists
	// TODO

	// Ensure the traefik plugin's network exists
	// TODO

	// Delete the existing service
	// TODO: Update and start the existing service if it already exists
	if existingService != nil {
		err = s.client.ServiceRemove(ctx, existingService.ID)
		if err != nil {
			return &server.Error{
				Code:    server.EINTERNAL,
				Message: "Failed to remove existing service",
				Op:      "docker.RuntimeService.Start",
				Err:     err,
			}
		}
	}

	// Create and start the service
	spec, err := s.getServiceSpec(ctx, app)
	if err != nil {
		return err
	}
	_, err = s.client.ServiceCreate(ctx, spec, types.ServiceCreateOptions{
		QueryRegistry: true,
	})
	return err
}

// Returns the existing service for the app or nil if it doesn't exist
func (s *RuntimeService) getExistingService(ctx context.Context, app internal.App, includeStatus bool) (*swarm.Service, error) {
	running, err := s.client.ServiceList(ctx, types.ServiceListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "name", Value: s.getServiceName(app)}),
		Status:  includeStatus,
	})
	if err != nil {
		return nil, err
	}
	s.logger.V("All running services: %v", running)

	for _, service := range running {
		if service.Spec.Annotations.Labels[miasmaIdLabel] == app.ID {
			return &service, nil
		}
	}
	return nil, nil
}

func (s *RuntimeService) getServiceName(app internal.App) string {
	return miasmaServiceNamePrefix + app.Name
}

// Convert an app into a docker service
func (s *RuntimeService) getServiceSpec(ctx context.Context, app internal.App) (swarm.ServiceSpec, error) {
	// TODO: get real environment
	env := map[string]string{}

	// Strip custom tag and use digest instead
	imageNoTag := app.Image
	if i := strings.LastIndex(imageNoTag, ":"); i >= 0 {
		imageNoTag = imageNoTag[0:i]
	}
	image := imageNoTag + "@" + app.ImageDigest

	labels := map[string]string{
		miasmaIdLabel: app.ID,
	}

	command := []string{}
	if app.Command != nil {
		command = append(command, *app.Command)
	}

	mounts := lo.Map(app.Volumes, func(volume *internal.BoundVolume, i int) mount.Mount {
		return mount.Mount{
			Source: volume.Source,
			Target: volume.Target,
			Type:   mount.TypeBind,
		}
	})

	envSlice := lo.Map(lo.Entries(env), func(entry lo.Entry[string, string], _ int) string {
		return fmt.Sprintf("%s=%s", entry.Key, fmt.Sprint(entry.Value))
	})

	return swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name:   s.getServiceName(app),
			Labels: labels,
		},
		TaskTemplate: swarm.TaskSpec{
			Placement: &swarm.Placement{
				Constraints: app.Placement,
			},
			ContainerSpec: &swarm.ContainerSpec{
				Image:   image,
				Env:     envSlice,
				Command: command,
				Mounts:  mounts,
			},
		},
	}, nil
}

// Stop implements server.RuntimeService
func (s *RuntimeService) Stop(ctx context.Context, app internal.App) error {
	s.logger.D("Stopping app: %s", app.Name)
	existing, err := s.getExistingService(ctx, app, false)
	if err != nil {
		return err
	}
	if existing == nil {
		// App already stopped
		s.logger.W("No existing service found, app already stopped")
		return nil
	}
	return s.client.ServiceRemove(ctx, existing.ID)
}

// ClusterInfo implements server.RuntimeService
func (s *RuntimeService) ClusterInfo(ctx context.Context) (*internal.ClusterInfo, error) {
	s.logger.D("Getting swarm details")
	info, err := s.client.Info(ctx)
	if err != nil {
		return nil, &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to run 'docker info'",
			Op:      "docker.RuntimeService.ClusterInfo()",
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
			Op:      "docker.RuntimeService.ClusterInfo()",
			Err:     err,
		}
	}
	return &internal.ClusterInfo{
		ID:          swarm.ID,
		JoinCommand: fmt.Sprintf("docker swarm join --token %s %s:2377", swarm.JoinTokens.Worker, info.Swarm.NodeAddr),
		CreatedAt:   swarm.CreatedAt,
		UpdatedAt:   swarm.UpdatedAt,
	}, nil
}

// Version implements server.RuntimeService
func (s *RuntimeService) Version(ctx context.Context) (string, error) {
	s.logger.D("Getting docker version")
	info, err := s.client.Info(ctx)
	return info.ServerVersion, err
}

// GetRuntimeAppInfo implements server.RuntimeService
func (s *RuntimeService) GetRuntimeAppInfo(ctx context.Context, app internal.App) (server.RuntimeAppInfo, error) {
	service, err := s.getExistingService(ctx, app, true)
	if err != nil {
		return EmptyRuntimeServiceDetails, err
	} else if service == nil {
		return EmptyRuntimeServiceDetails, &server.Error{
			Code:    server.ENOTFOUND,
			Message: "No running service found",
			Op:      "docker.RuntimeService.GetService",
		}
	}

	return server.RuntimeAppInfo{
		Instances: internal.AppInstances{
			Total:   int32(service.ServiceStatus.DesiredTasks),
			Running: int32(service.ServiceStatus.RunningTasks),
		},
		Status: "running",
	}, nil
}

// RestartRunningApps implements server.RuntimeService
func (s *RuntimeService) RestartRunningApps(ctx context.Context, apps []internal.App) error {
	all, err := s.client.ServiceList(ctx, types.ServiceListOptions{
		Status: true,
	})
	if err != nil {
		return err
	}

	for _, service := range all {
		isRunning := service.ServiceStatus != nil && service.ServiceStatus.DesiredTasks >= 0
		app, hasApp := lo.Find(apps, func(app internal.App) bool {
			return service.Spec.Annotations.Labels[miasmaIdLabel] == app.ID
		})
		if isRunning && hasApp {
			err = s.Restart(ctx, app)
			if err != nil {
				s.logger.W("Failed to restart app '%s': %v", app.Name, err)
			}
		}
	}

	return nil
}
