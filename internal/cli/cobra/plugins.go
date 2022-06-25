package cobra

import (
	"context"
	"fmt"
	"strings"

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
	ctx := context.Background()
	title.Println("\nPlugins")

	plugins, err := api.ListPlugins(ctx, `{ name enabled }`)
	checkErr(err)

	for _, p := range plugins {
		status := red("●")
		if p.Enabled {
			status = green("●")
		}
		listItem := []string{status, string(p.Name)}
		if !p.Enabled {
			listItem = append(listItem, dim("(disabled)"))
		}
		fmt.Printf(" %s\n", strings.Join(listItem, " "))
	}

	fmt.Println()
}
