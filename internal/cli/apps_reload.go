package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsReloadCmd = &cobra.Command{
	Use:   "apps:reload",
	Short: "Reload an application or start it if it's not already running",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		reloadApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsReloadCmd)
	flags.UseAppFlag(appsReloadCmd)
}

func reloadApp(appName string) {
	fmt.Printf("Reloading %s...\n", appName)
	reloadParams := operations.NewReloadAppParams().WithAppName(appName)
	_, err := config.Client().Operations.ReloadApp(reloadParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
