package cli

import (
	"github.com/spf13/cobra"
)

var pluginRemoveCmd = &cobra.Command{
	Use:   "plugins:remove",
	Short: "Stop and remove an installed plugin",
	Long: `Stop and remove an installed plugins: PostgreSQL, Mongo
	
Plugins are simple, pre-configured, reusable applications. Every plugin can also be defined as an
app.`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"postgres", "mongo"},
	Run: func(cmd *cobra.Command, args []string) {
		removePlugin(args[0])
	},
}

func init() {
	RootCmd.AddCommand(pluginRemoveCmd)
}

func removePlugin(pluginName string) {
	panic("NOT IMPLEMENTED")
	// Stop
	// Remove
}
