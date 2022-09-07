package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/server"
	"github.com/aklinker1/miasma/internal/server/zero"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

type runtimeServiceRepo struct {
	client           client.APIClient
	logger           server.Logger
	certResolverName string
}

func NewRuntimeServiceRepo(logger server.Logger, client *client.Client, certResolverName string) server.RuntimeServiceRepo {
	return &runtimeServiceRepo{
		client:           client,
		logger:           logger,
		certResolverName: certResolverName,
	}
}

// Start implements server.RuntimeServiceRepo
func (s *runtimeServiceRepo) Create(ctx context.Context, service server.RuntimeServiceSpec) error {
	s.logger.D("Creating service: %s", service.App.Name)
	dockerSpec, err := s.getDockerSpec(ctx, service)
	if err != nil {
		return err
	}

	// Ensure the network exists for intra-app communication
	err = s.ensureNetwork(ctx, defaultNetwork)
	if err != nil {
		return err
	}

	// Create (and start) a new service
	_, err = s.client.ServiceCreate(ctx, dockerSpec, types.ServiceCreateOptions{
		QueryRegistry: true,
	})
	if err != nil {
		return &server.Error{
			Code:    server.EINTERNAL,
			Message: "Failed to create service",
			Op:      "docker.runtimeServiceRepo.Start",
			Err:     err,
		}
	}

	return nil
}

// Returns the existing service for the app or nil if it doesn't exist
func (s *runtimeServiceRepo) GetAll(ctx context.Context, filter server.RuntimeServicesFilter) ([]server.RuntimeService, error) {
	search := []filters.KeyValuePair{{
		Key:   "label",
		Value: fmt.Sprintf("%s=true", miasmaFlagLabel),
	}}
	if filter.ID != nil {
		search = append(search, filters.KeyValuePair{
			Key:   "id",
			Value: *filter.ID,
		})
	}
	if filter.AppID != nil {
		search = append(search, filters.KeyValuePair{
			Key:   "label",
			Value: fmt.Sprintf("%s=%s", miasmaIdLabel, *filter.AppID),
		})
	}
	services, err := s.client.ServiceList(ctx, types.ServiceListOptions{
		Filters: filters.NewArgs(search...),
		Status:  filter.IncludeStatus,
	})
	if err != nil {
		return nil, err
	}
	return lo.Map(services, func(s swarm.Service, _ int) server.RuntimeService {
		return server.RuntimeService{
			Service: s,
			AppID:   s.Spec.Labels[miasmaIdLabel],
		}
	}), nil
}

// GetRuntimeAppInfo implements server.RuntimeServiceRepo
func (s *runtimeServiceRepo) GetOne(ctx context.Context, filter server.RuntimeServicesFilter) (server.RuntimeService, error) {
	services, err := s.GetAll(ctx, filter)
	if err != nil {
		return zero.RuntimeService, err
	}
	if len(services) == 0 {
		return zero.RuntimeService, &server.Error{
			Code:    server.ENOTFOUND,
			Message: fmt.Sprintf("Service not found for %+v", filter),
		}
	}
	return services[0], nil
}

func (s *runtimeServiceRepo) getOneOrNil(ctx context.Context, filter server.RuntimeServicesFilter) (*server.RuntimeService, error) {
	services, err := s.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, nil
	}
	return &services[0], nil
}

// getServiceName returns a valid DNS target name based on the app name
func (s *runtimeServiceRepo) getServiceName(appName string) string {
	return strings.ReplaceAll(strings.ToLower(appName), " ", "-")
}

func (s *runtimeServiceRepo) getNetworkName(appName string) string {
	return miasmaNetworkNamePrefix + appName
}

// Convert an app into a docker service
func (s *runtimeServiceRepo) getDockerSpec(ctx context.Context, service server.RuntimeServiceSpec) (swarm.ServiceSpec, error) {
	app := service.App
	readonlyEnv := service.Env
	plugins := service.Plugins
	route := service.Route

	hostname := s.getServiceName(app.Name)

	env := map[string]string{}
	if readonlyEnv != nil {
		for key, value := range readonlyEnv {
			env[key] = value
		}
	}

	// Strip custom tag and use digest instead
	imageNoTag := app.Image
	if i := strings.LastIndex(imageNoTag, ":"); i >= 0 {
		imageNoTag = imageNoTag[0:i]
	}
	image := imageNoTag + "@" + app.ImageDigest

	command := app.Command

	ports, err := s.getPorts(ctx, app)
	if err != nil {
		return swarm.ServiceSpec{}, err
	}
	env["PORT"] = fmt.Sprint(ports[0].TargetPort)
	for i := 0; i < len(ports); i++ {
		env[fmt.Sprintf("PORT_%d", i+1)] = fmt.Sprint(ports[i].TargetPort)
	}

	labels := map[string]string{
		miasmaIdLabel:   app.ID,
		miasmaFlagLabel: "true",
	}

	traefikPlugin, ok := lo.Find(plugins, func(plugin internal.Plugin) bool {
		return plugin.Name == internal.PluginNameTraefik
	})
	if ok && traefikPlugin.Enabled && route != nil {
		labels["traefik.enable"] = "true"
		labels["traefik.docker.network"] = s.getNetworkName(defaultNetwork)
		labels["traefik.http.services."+hostname+".loadbalancer.server.port"] = fmt.Sprint(ports[0].TargetPort)

		ruleLabel := "traefik.http.routers." + hostname + ".rule"
		if route.TraefikRule != nil {
			labels[ruleLabel] = *route.TraefikRule
		} else if route.Host != nil && route.Path != nil {
			labels[ruleLabel] = fmt.Sprintf("(Host(`%s`) && PathPrefix(`%s`))", *route.Host, *route.Path)
		} else if route.Host != nil {
			labels[ruleLabel] = fmt.Sprintf("Host(`%s`)", *route.Host)
		}

		// HTTPS
		traefikConfig := traefikPlugin.ConfigForTraefik()
		if traefikConfig.EnableHttps {
			tlsLabel := fmt.Sprintf("traefik.http.routers.%s.tls", hostname)
			labels[tlsLabel] = "true"
			tlsResolverLabel := fmt.Sprintf("traefik.http.routers.%s.tls.certresolver", hostname)
			labels[tlsResolverLabel] = s.certResolverName
		}
	}

	mounts := lo.Map(app.Volumes, func(volume *internal.BoundVolume, i int) mount.Mount {
		return mount.Mount{
			Source: volume.Source,
			Target: volume.Target,
			Type:   mount.TypeBind,
		}
	})

	networks := lo.Map(app.Networks, func(networkName string, _ int) swarm.NetworkAttachmentConfig {
		return swarm.NetworkAttachmentConfig{
			Target: networkName, // Don't use s.getNetworkName here, these are specified by the user
		}
	})
	networks = append(networks, swarm.NetworkAttachmentConfig{
		Target: s.getNetworkName(defaultNetwork),
	})

	for key := range env {
		if !dockerEnvKeyRegex.Match([]byte(key)) {
			return zero.SwarmServiceSpec, &server.Error{
				Code: server.EINVALID,
				Message: fmt.Sprintf(
					"Docker environment variables must match /%s/, but '%s' did not",
					dockerEnvKeyRegex.String(),
					key,
				),
				Op: "docker.runtimeServiceRepo.getDockerSpec",
			}
		}
	}
	envSlice := lo.Map(lo.Entries(env), func(entry lo.Entry[string, string], _ int) string {
		return fmt.Sprintf("%s=%s", entry.Key, fmt.Sprint(entry.Value))
	})

	return swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name:   hostname,
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
			Networks: networks,
		},
		EndpointSpec: &swarm.EndpointSpec{
			Ports: ports,
		},
	}, nil
}

// Stop implements server.RuntimeServiceRepo
func (s *runtimeServiceRepo) Remove(ctx context.Context, service server.RuntimeService) (server.RuntimeService, error) {
	s.logger.D("Removing service: %s", service.Spec.Name)
	existing, err := s.getOneOrNil(ctx, server.RuntimeServicesFilter{
		ID: &service.ID,
	})
	if err != nil {
		return zero.RuntimeService, err
	}
	if existing == nil {
		// App already stopped
		s.logger.W("No existing service found, app already stopped")
		return zero.RuntimeService, nil
	}
	return *existing, s.client.ServiceRemove(ctx, existing.ID)
}

func (s *runtimeServiceRepo) ensureNetwork(ctx context.Context, networkName string) error {
	s.logger.D("Ensuring network exists: %s", s.getNetworkName(networkName))
	networks, err := s.client.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key: "name", Value: s.getNetworkName(networkName),
		}),
	})
	if err != nil {
		return err
	}

	s.logger.V("Queried networks: %+v", networks)
	if len(networks) > 0 {
		if networks[0].Driver == "overlay" && networks[0].Scope == "swarm" && networks[0].Labels[miasmaFlagLabel] == "true" {
			s.logger.V("Network already exists and is configured correctly")
			return nil
		}
		err = s.client.NetworkRemove(ctx, networks[0].ID)
		if err != nil {
			return err
		}
	}

	_, err = s.client.NetworkCreate(ctx, s.getNetworkName(networkName), types.NetworkCreate{
		Driver: "overlay",
		Scope:  "swarm",
		Labels: map[string]string{
			miasmaFlagLabel: "true",
		},
	})
	return err
}

func (s *runtimeServiceRepo) getPorts(ctx context.Context, app internal.App) ([]swarm.PortConfig, error) {
	required := len(app.TargetPorts)
	if len(app.PublishedPorts) > required {
		required = len(app.PublishedPorts)
	}
	if required == 0 {
		required = 1
	}

	target := app.TargetPorts
	for len(target) < required {
		target = append(target, utils.RandInt(3000, 4000))
	}

	published := app.PublishedPorts
	if required != len(published) {
		openPorts, err := s.findOpenPorts(ctx, required-len(published))
		if err != nil {
			return nil, err
		}
		published = append(published, openPorts...)
	}

	ports := []swarm.PortConfig{}
	for i := 0; i < required; i++ {
		ports = append(ports, swarm.PortConfig{
			PublishedPort: uint32(published[i]),
			TargetPort:    uint32(target[i]),
		})
	}

	s.logger.V("Target ports: %+v", target)
	s.logger.V("Published ports: %+v", published)
	return ports, nil
}

func (s *runtimeServiceRepo) findOpenPorts(ctx context.Context, count int) ([]int, error) {
	s.logger.D("Finding %d open port%s", count, lo.Ternary(count == 1, "", "s"))
	services, err := s.client.ServiceList(ctx, types.ServiceListOptions{})
	if err != nil {
		return nil, err
	}
	filledPorts := map[int]bool{}
	for _, service := range services {
		for _, port := range service.Endpoint.Ports {
			filledPorts[int(port.PublishedPort)] = true
		}
	}
	results := []int{}
	for port := 3001; port < 4000 && len(results) < count; port++ {
		if _, ok := filledPorts[port]; !ok {
			results = append(results, port)
		}
	}
	if len(results) < count {
		return nil, fmt.Errorf("Not enough available ports to start the service (required=%d, available=%d)", count, len(results))
	}
	return results, nil
}

func (s *runtimeServiceRepo) Update(ctx context.Context, serviceID string, newService server.RuntimeServiceSpec) (server.RuntimeService, error) {
	newDockerSpec, err := s.getDockerSpec(ctx, newService)
	if err != nil {
		return zero.RuntimeService, err
	}
	swarm, err := s.client.SwarmInspect(ctx)
	if err != nil {
		return zero.RuntimeService, err
	}
	s.client.ServiceUpdate(ctx, serviceID, swarm.Version, newDockerSpec, types.ServiceUpdateOptions{})
	return zero.RuntimeService, server.NewNotImplementedError("docker.runtimeServiceRepo.Update")
}
