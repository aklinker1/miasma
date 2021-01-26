package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsStartCmd = &cobra.Command{
	Use:   "apps:start",
	Short: "Start an application",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		startApp(flags.GetAppFlag(cmd))
	},
}

func init() {
	RootCmd.AddCommand(appsStartCmd)
	flags.UseAppFlag(appsStartCmd)
}

func startApp(appName string) {
	fmt.Printf("Starting %s...\n", appName)
	startParams := operations.NewStartAppParams()
	startParams.AppName = appName
	_, err := config.Client().Operations.StartApp(startParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
