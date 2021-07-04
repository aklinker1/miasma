package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var traefikCmd = &cobra.Command{
	Use:   "traefik",
	Short: "Get routing rules for an app",
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		getRouting(appName)
	},
}

func init() {
	RootCmd.AddCommand(traefikCmd)
	flags.UseAppFlag(traefikCmd)
}

func getRouting(appName string) {
	fmt.Printf("Getting Traefik config for %s...\n", appName)
	client := config.Client()
	app, err := client.Operations.GetApp(operations.NewGetAppParams().WithAppName(appName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config, err := client.Operations.GetAppTraefikConfig(operations.NewGetAppTraefikConfigParams().WithAppID(app.Payload.ID))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config.Payload.Host != nil {
		fmt.Printf("Host: %v\n", *config.Payload.Host)
	}
	if config.Payload.Path != nil {
		fmt.Printf("Path: %v\n", *config.Payload.Path)
	}
	if config.Payload.TraefikRule != nil {
		fmt.Printf("Rule: %v\n", *config.Payload.TraefikRule)
	}
	fmt.Println("Done!")
}
