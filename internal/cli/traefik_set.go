package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/aklinker1/miasma/package/models"
	"github.com/spf13/cobra"
)

var traefikSetCmd = &cobra.Command{
	Use:   "traefik:set",
	Short: "Set routing rules for an app",
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()
		routing := flags.GetSetTraefikFlags(cmd)

		setRouting(appName, routing)
	},
}

func init() {
	RootCmd.AddCommand(traefikSetCmd)
	flags.UseAppFlag(traefikSetCmd)
	flags.UseSetTraefikFlags(traefikSetCmd)
}

func setRouting(appName string, routing *flags.SetTraefik) {
	client := config.Client()
	app, err := client.Operations.GetApp(operations.NewGetAppParams().WithAppName(appName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client.Operations.UpdateAppTraefikConfig(
		operations.NewUpdateAppTraefikConfigParams().
			WithAppID(app.Payload.ID).
			WithNewTraefikConfig(&models.InputTraefikPluginConfig{
				Host:        routing.Host,
				Path:        routing.Path,
				TraefikRule: routing.Rule,
			}),
	)
}
