package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List apps",
	Run: func(cmd *cobra.Command, args []string) {
		listApps(flags.GetAllFlag(cmd))
	},
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flags.UseAllFlag(appsCmd)
}

func listApps(includeHidden bool) {
	fmt.Println("List apps:")
	client := config.Client()
	apps, err := client.Operations.GetApps(operations.NewGetAppsParams().WithHidden(&includeHidden))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, app := range apps.Payload {
		fmt.Printf(" - %s\n", *app.Name)
	}
	fmt.Printf("(%d total)\n", len(apps.Payload))
}
