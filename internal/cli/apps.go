package cli

import (
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List apps",
	Run: func(cmd *cobra.Command, args []string) {
		listApps(flags.GetAllFlag(cmd))
	},
}

func init() {
	RootCmd.AddCommand(appsCmd)
	flags.UseAllFlag(appsCmd)
}

func listApps(includeHidden bool) {
	panic("NOT IMPLEMENTED")
}
