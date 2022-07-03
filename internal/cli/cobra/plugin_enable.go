package cobra

import (
	"context"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var pluginEnableCmd = &cobra.Command{
	Use:       "plugins:enable",
	Short:     "Enable and start a pre-defined plugin",
	Long:      `Enable one of the pre-defined plugins. Plugins are simple, pre-configured, reusable applications that have custom integrations with Miasma.`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"TRAEFIK"},
	Run: func(cmd *cobra.Command, args []string) {
		pluginConfig := flags.GetPluginConfigFlag(cmd)
		pluginName := args[0]
		enablePlugin(pluginName, pluginConfig)
	},
}

func init() {
	RootCmd.AddCommand(pluginEnableCmd)
	flags.UsePluginConfigFlag(pluginEnableCmd)
}

func enablePlugin(pluginName string, pluginConfig map[string]any) {
	ctx := context.Background()
	title.Printf("\nEnabling %s...\n", pluginName)

	err := api.EnablePlugin(ctx, pluginName, pluginConfig)
	checkErr(err)

	done("%s enabled", pluginName)
}
