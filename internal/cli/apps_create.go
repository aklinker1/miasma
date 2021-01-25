package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/internal/cli/validation"
	"github.com/spf13/cobra"
)

var appsCreateCmd = &cobra.Command{
	Use:   "apps:create",
	Short: "Create and deploy a new application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName, err := validation.AppName(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		image := flags.GetImageFlag(cmd)
		hidden := flags.GetHiddenFlag(cmd)

		createApp(appName, image, hidden)
	},
}

func init() {
	RootCmd.AddCommand(appsCreateCmd)
	flags.UseImageFlag(appsCreateCmd)
	flags.UseHiddenFlag(appsCreateCmd)
}

func createApp(appName string, image string, hidden bool) {
	panic("NOT IMPLEMENTED")
	// Create
	// Start
}
