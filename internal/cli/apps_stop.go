package cli

import (
	"github.com/aklinker1/miasma/internal/cli/flags"
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
	panic("NOT IMPLEMENTED")
	// Stop
}
