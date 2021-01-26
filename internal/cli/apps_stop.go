package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsStopCmd = &cobra.Command{
	Use:   "apps:stop",
	Short: "Stop a running application",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		stopApp(flags.GetAppFlag(cmd))
	},
}

func init() {
	RootCmd.AddCommand(appsStopCmd)
	flags.UseAppFlag(appsStopCmd)
}

func stopApp(appName string) {
	fmt.Printf("Stopping %s...\n", appName)
	stopParams := operations.NewStopAppParams()
	stopParams.AppName = appName
	_, err := config.Client().Operations.StopApp(stopParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
