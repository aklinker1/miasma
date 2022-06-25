package cobra

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsDeleteCmd = &cobra.Command{
	Use:   "apps:delete",
	Short: "Delete an existing application",
	Long:  `Delete an existing application. If the app is running, it is stopped first`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		deleteApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsDeleteCmd)
	flags.UseAppFlag(appsDeleteCmd)
}

func deleteApp(appName string) {
	ctx := context.Background()
	title.Printf("\nDeleting %s...\n", appName)

	app, err := api.GetApp(ctx, appName, `{ id name }`)
	checkErr(err)

	err = api.DeleteApp(ctx, app.ID)
	checkErr(err)

	fmt.Printf("Done!\n\n")
}
