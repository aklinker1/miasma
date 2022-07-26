package docker

import "regexp"

var dockerEnvKeyRegex = regexp.MustCompile("^[0-9A-Z_]+$")

var (
	miasmaIdLabel           = "miasma-id"
	miasmaFlagLabel         = "miasma"
	miasmaNetworkNamePrefix = "miasma-"
	defaultNetwork          = "default"
)
