package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var pluginRemoveCmd = &cobra.Command{
	Use:   "plugins:remove",
	Short: "Stop and remove an installed plugin",
	Long: `Stop and remove an installed plugins: Traefik
	
Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"traefik"},
	Run: func(cmd *cobra.Command, args []string) {
		removePlugin(args[0])
	},
}

func init() {
	RootCmd.AddCommand(pluginRemoveCmd)
}

func removePlugin(pluginName string) {
	fmt.Printf("Removing %s...\n", pluginName)
	uninstallParams := operations.NewUninstallPluginParams()
	uninstallParams.PluginName = pluginName
	_, err := config.Client().Operations.UninstallPlugin(uninstallParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
