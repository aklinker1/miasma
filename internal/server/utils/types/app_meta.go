package types

type Route struct {
	Host        string
	Path        string
	TraefikRule string `yaml:"traefikRule"`
}

type AppMetaData struct {
	// App
	Name   string
	Image  *string
	Hidden *bool

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
