package cli

import (
	"github.com/spf13/cobra"
)

var pluginInstallCmd = &cobra.Command{
	Use:   "plugins:install",
	Short: "Install and start a pre-defined plugin",
	Long: `Install one of the pre-defined plugins: PostgreSQL, Mongo
	
Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"postgres", "mongo"},
	Run: func(cmd *cobra.Command, args []string) {
		installPlugin(args[0])
	},
}

func init() {
	RootCmd.AddCommand(pluginInstallCmd)
}

func installPlugin(pluginName string) {
	panic("NOT IMPLEMENTED")
	// Install
	// Start
}
