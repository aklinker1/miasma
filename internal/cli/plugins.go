package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/package/client/operations"
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
	fmt.Println("Available plugins:")
	client := config.Client()
	plugins, err := client.Operations.ListPlugins(operations.NewListPluginsParams())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, app := range plugins.Payload {
		installed := ""
		if app.Installed {
			installed = " (installed)"
		}
		fmt.Printf(" - %s%s\n", app.Name, installed)
	}
	fmt.Printf("(%d total)\n", len(plugins.Payload))
}
