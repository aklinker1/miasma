package cli

import (
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "List plugins",
	Run: func(cmd *cobra.Command, args []string) {
		listPlugins()
	},
}

func init() {
	RootCmd.AddCommand(pluginsCmd)
}

func listPlugins() {
	panic("NOT IMPLEMENTED")
}
