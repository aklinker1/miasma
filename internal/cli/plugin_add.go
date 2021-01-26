package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var pluginInstallCmd = &cobra.Command{
	Use:   "plugins:add",
	Short: "Install and start a pre-defined plugin",
	Long: `Install one of the pre-defined plugins: Traefik
	
Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"traefik"},
	Run: func(cmd *cobra.Command, args []string) {
		installPlugin(args[0])
	},
}

func init() {
	RootCmd.AddCommand(pluginInstallCmd)
}

func installPlugin(pluginName string) {
	fmt.Printf("Adding %s...\n", pluginName)
	installParams := operations.NewInstallPluginParams()
	installParams.PluginName = pluginName
	_, err := config.Client().Operations.InstallPlugin(installParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
