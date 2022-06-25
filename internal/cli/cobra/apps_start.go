package cobra

import (
	"context"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsStartCmd = &cobra.Command{
	Use:   "apps:start",
	Short: "Start an application",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		startApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsStartCmd)
	flags.UseAppFlag(appsStartCmd)
}

func startApp(appName string) {
	ctx := context.Background()
	title.Printf("\nStarting %s...\n", appName)

	app, err := api.GetApp(ctx, appName, `{ id name }`)
	checkErr(err)

	err = api.StartApp(ctx, app.ID)
	checkErr(err)

	done("%s started", appName)
}
