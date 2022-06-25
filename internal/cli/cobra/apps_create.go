package cobra

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsCreateCmd = &cobra.Command{
	Use:   "apps:create",
	Short: "Create and deploy a new application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
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
	ctx := context.Background()

	title.Printf("\nCreating %s...\n", appName)

	if image == "" {
		panic("An image is required, pass with '--image some/name'")
	}

	err := api.CreateApp(ctx, internal.AppInput{
		Name:   appName,
		Image:  image,
		Hidden: &hidden,
	})
	checkErr(err)

	fmt.Printf("%s started\n\n", appName)
}
