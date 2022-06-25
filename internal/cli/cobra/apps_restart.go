package cobra

import (
	"context"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsRestartCmd = &cobra.Command{
	Use:   "apps:restart",
	Short: "Restart an application or start it if it's not already running",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		reloadApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsRestartCmd)
	flags.UseAppFlag(appsRestartCmd)
}

func reloadApp(appName string) {
	ctx := context.Background()
	title.Printf("\nRestarting %s...\n", appName)

	app, err := api.GetApp(ctx, appName, `{ id name }`)
	checkErr(err)

	err = api.RestartApp(ctx, app.ID)
	checkErr(err)

	done("%s restarted", appName)
}
