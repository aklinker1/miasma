package constants

import (
	"fmt"

	"github.com/aklinker1/miasma/internal/server/utils/env"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared"
)

type predefinedPlugins struct {
	Traefik  types.AppMetaData `json:"traefik"`
	Postgres types.AppMetaData `json:"postgres"`
	Mongo    types.AppMetaData `json:"mongo"`
	// Redis    types.AppMetaData `json:"redis"`
}

var Plugins = predefinedPlugins{
	// https://doc.traefik.io/traefik/getting-started/quick-start/
	Traefik: types.AppMetaData{
		Name:           "plugin-traefik",
		Image:          shared.StringPtr("traefik:v2.4"),
		TargetPorts:    []uint32{80, 8080},
		PublishedPorts: []uint32{80, 4000},
		Placement:      []string{"node.labels.traefik==true"},
		Volumes: []string{
			"/var/run/docker.sock:/var/run/docker.sock",
		},
		Command: []string{"traefik", "--api.insecure=true", "--providers.docker", "--providers.docker.swarmmode"},
	},

	Postgres: types.AppMetaData{
		Name:           "plugin-postgres",
		Image:          shared.StringPtr("postgres:alpine"),
		TargetPorts:    []uint32{5432},
		PublishedPorts: []uint32{5432},
		Command:        []string{"-p 4001"},
		Placement:      []string{"node.labels.postgres==true"},
		Volumes:        []string{fmt.Sprintf("%s:/var/lib/postgresql/data", env.PLUGIN_POSTGRES_DATA_VOLUME)},
		Env: map[string]string{
			"POSTGRES_PASSWORD": env.PLUGIN_POSTGRES_PASSWORD,
			"POSTGRES_USER":     env.PLUGIN_POSTGRES_USER,
			"POSTGRES_DB":       "miasma",
		},
	},

	Mongo: types.AppMetaData{
		Name:           "plugin-mongo",
		Image:          shared.StringPtr("mongo:4.4"),
		TargetPorts:    []uint32{27017},
		PublishedPorts: []uint32{4003},
		Command:        []string{"--bind_ip_all"},
		Placement:      []string{"node.labels.mongo==true"},
		Volumes:        []string{fmt.Sprintf("%s:/data/db", env.PLUGIN_MONGO_DATA_VOLUME)},
	},
}
