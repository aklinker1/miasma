package cobra

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsStopCmd = &cobra.Command{
	Use:   "apps:stop",
	Short: "Stop a running application",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		stopApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsStopCmd)
	flags.UseAppFlag(appsStopCmd)
}

func stopApp(appName string) {
	ctx := context.Background()
	title.Printf("\nStopping %s...\n", appName)

	app, err := api.GetApp(ctx, appName, `{ id name }`)
	checkErr(err)

	err = api.StopApp(ctx, app.ID)
	checkErr(err)

	fmt.Printf("Done!\n\n")
}
