package cobra

import (
	"context"

	"github.com/spf13/cobra"
)

var pluginDisableCmd = &cobra.Command{
	Use:       "plugins:disable",
	Short:     "Stop and disable a plugin",
	Long:      `Stop and disable a plugins. Plugins are simple, pre-configured, reusable applications that have custom integrations with Miasma.`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"TRAEFIK"},
	Run: func(cmd *cobra.Command, args []string) {
		disablePlugin(args[0])
	},
}

func init() {
	RootCmd.AddCommand(pluginDisableCmd)
}

func disablePlugin(pluginName string) {
	ctx := context.Background()
	title.Printf("\nDisabling %s...\n", pluginName)

	err := api.DisablePlugin(ctx, pluginName)
	checkErr(err)

	done("%s disabled", pluginName)
}
