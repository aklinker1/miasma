package types

import "github.com/aklinker1/miasma/internal/server/gen/models"

type Route struct {
	Host        *string
	Path        *string
	TraefikRule *string `yaml:"traefikRule"`
}

type AppMetaDataWithoutName struct {
	// App
	Image  string
	Hidden bool

	// Config
	TargetPorts    []uint32 `yaml:"targetPorts"`
	PublishedPorts []uint32 `yaml:"publishedPorts"`
	Placement      []string
	Networks       []string
	Volumes        []*models.AppConfigVolumesItems0
	Command        []string
	Route          *Route

	// Env
	Env map[string]interface{}
}

type AppMetaData struct {
	AppMetaDataWithoutName
	Name string
}
