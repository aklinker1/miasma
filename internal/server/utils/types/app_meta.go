package types

type AppMetaData struct {
	// App
	Name   string
	Image  *string
	Hidden *bool

	// Config
	TargetPorts []uint32
	Networks    []string
	Plugins     []string

	// Env
	Env map[string]string
}
