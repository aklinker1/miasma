package app_mapper

import (
	"fmt"
	"strings"

	"docker.io/go-docker/api/types/mount"
	"docker.io/go-docker/api/types/swarm"
	"github.com/aklinker1/miasma/internal/server/utils/constants"
	"github.com/aklinker1/miasma/internal/server/utils/mappers/volume_mapper"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func ToService(details *server_models.AppDetails, customEnv map[string]string, plugins *server_models.AppPlugins, getNextPorts func(int) ([]uint32, error), digest string) (*swarm.ServiceSpec, error) {
	// Setup ports
	var targetPorts = []uint32{}
	if len(details.RunConfig.TargetPorts) > 0 {
		targetPorts = append(targetPorts, details.RunConfig.TargetPorts...)
	} else {
		// Internal to container. This does not have the same limitations (3001-3999) at the
		// published ports. The range being close is a coincidence
		targetPorts = append(targetPorts, shared.RandUInt32(3000, 4000))
	}
	envPorts := []string{fmt.Sprintf("PORT=%d", targetPorts[0])}
	publishedPorts, err := getNextPorts(len(targetPorts))
	if err != nil {
		return nil, err
	}
	portConfigs := []swarm.PortConfig{}
	for index, targetPort := range targetPorts {
		publishedPort := publishedPorts[index]
		envPort := fmt.Sprintf("PORT_%d=%d", index+1, targetPort)
		envPorts = append(envPorts, envPort)
		portConfigs = append(portConfigs, swarm.PortConfig{
			PublishedPort: publishedPort,
			TargetPort:    targetPort,
		})
		if index == 0 {
			log.V("Ports Env: %d:%d (%s, %s)", publishedPort, targetPort, envPorts[0], envPort)
		} else {
			log.V("Ports Env: %d:%d (%s)", publishedPort, targetPort, envPort)
		}
	}

	// Setup networks
	networks := []swarm.NetworkAttachmentConfig{
		{
			Target: details.App.Name,
		},
	}
	for _, network := range details.RunConfig.Networks {
		log.V("Connecting to another app: %s", network)
		networks = append(networks, swarm.NetworkAttachmentConfig{
			Target: network,
		})
	}
	if plugins.Traefik != nil {
		log.V("Adding traefik network")
		networks = append(networks, swarm.NetworkAttachmentConfig{
			Target: constants.Plugins.Traefik.App.Name,
		})
	}
	// TODO: More plugins
	// if plugins.Postgres {
	// 	log.V("Adding postgres network")
	// 	networks = append(networks, swarm.NetworkAttachmentConfig{
	// 		Target: constants.Plugins.Postgres.Name,
	// 	})
	// }
	// if plugins.Mongo {
	// 	log.V("Adding mongo network")
	// 	networks = append(networks, swarm.NetworkAttachmentConfig{
	// 		Target: constants.Plugins.Mongo.Name,
	// 	})
	// }

	// Setup env variables
	env := append(envPorts, []string{}...)
	for key, value := range customEnv {
		env = append(env, fmt.Sprintf("%s=%s", key, fmt.Sprint(value)))
		log.V("Added env: %s", key)
	}

	// Setup volumes
	volumes := []mount.Mount{}
	for _, volume := range details.RunConfig.Volumes {
		log.V("Added volume: %v", volume)
		volumes = append(volumes, volume_mapper.ToDocker(volume))
	}

	// Setup labels
	labels := map[string]string{
		"miasma": "true",
	}
	if plugins.Traefik != nil {
		enabled := "traefik.enable"
		networkLabel := "traefik.docker.network"
		targetPort := fmt.Sprintf("traefik.http.services.%s-service.loadbalancer.server.port", details.App.Name)
		rulesLabel := fmt.Sprintf("traefik.http.routers.%s.rule", details.App.Name)

		labels[enabled] = "true"
		labels[networkLabel] = constants.Plugins.Traefik.App.Name
		labels[targetPort] = fmt.Sprint(targetPorts[0])
		if plugins.Traefik.TraefikRule != nil {
			labels[rulesLabel] = *plugins.Traefik.TraefikRule
		} else if plugins.Traefik.Path == nil {
			labels[rulesLabel] = fmt.Sprintf("Host(`%s`)", *plugins.Traefik.Host)
		} else {
			labels[rulesLabel] = fmt.Sprintf("(Host(`%s`) && PathPrefix(`%s`))", *plugins.Traefik.Host, *plugins.Traefik.Path)
		}
	}
	var imageRepo string
	if strings.ContainsRune(details.App.Image, ':') {
		imageRepo = details.App.Image[0:strings.LastIndex(details.App.Image, ":")]
	} else {
		imageRepo = details.App.Image
	}
	image := imageRepo + "@" + digest
	log.V("Service will use '%s' instead of '%s'", image, details.App.Image)

	return &swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name:   details.App.Name,
			Labels: labels,
		},
		TaskTemplate: swarm.TaskSpec{
			Placement: &swarm.Placement{
				Constraints: details.RunConfig.Placement,
			},
			ContainerSpec: &swarm.ContainerSpec{
				Image:   image,
				Env:     env,
				Command: details.RunConfig.Command,
				Mounts:  volumes,
			},
			Networks: networks,
		},
		EndpointSpec: &swarm.EndpointSpec{
			Ports: portConfigs,
		},
	}, nil
}
