package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsDestroyCmd = &cobra.Command{
	Use:   "apps:destroy",
	Short: "Destroy an existing application",
	Long:  `Destroy an existing application. If the app is running, it is stopped first`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		destroyApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsDestroyCmd)
	flags.UseAppFlag(appsDestroyCmd)
}

func destroyApp(appName string) {
	fmt.Printf("Destrying %s...\n", appName)
	client := config.Client()
	deleteParams := operations.NewDeleteAppParams()
	deleteParams.AppName = appName
	_, err := client.Operations.DeleteApp(deleteParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done")
}
