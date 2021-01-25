package cli

import (
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsDestroyCmd = &cobra.Command{
	Use:   "apps:destroy",
	Short: "Destroy an existing application",
	Long:  `Destroy an existing application. If the app is running, it is stopped first`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		destroyApp(flags.GetAppFlag(cmd))
	},
}

func init() {
	RootCmd.AddCommand(appsDestroyCmd)
	flags.UseAppFlag(appsDestroyCmd)
}

func destroyApp(appName string) {
	panic("NOT IMPLEMENTED")
	// Destroy
	// Start
}
