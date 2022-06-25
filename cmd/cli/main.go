package main

import (
	"github.com/aklinker1/miasma/internal/cli"
	"github.com/aklinker1/miasma/internal/cli/cobra"
	"github.com/aklinker1/miasma/internal/cli/http"
)

// Compile time variables
var (
	VERSION    string
	BUILD      string
	BUILD_HASH string
	BUILD_DATE string
)

func main() {
	metadata := cli.Metadata{
		Version:   VERSION,
		Build:     BUILD,
		BuildHash: BUILD_HASH,
		BuildDate: BUILD_DATE,
	}
	api := http.DefaultMiasmaAPIClient()

	cobra.Execute(metadata, api)
}
