package cli

import (
	"github.com/aklinker1/miasma/internal/cli/flags"
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
	panic("NOT IMPLEMENTED")
	// Start
}
