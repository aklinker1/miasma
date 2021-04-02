package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsUpdateCmd = &cobra.Command{
	Use:   "apps:update",
	Short: "Create and deploy a new application",
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()
		image := flags.GetNewImageFlag(cmd)

		updateApp(appName, image)
	},
}

func init() {
	RootCmd.AddCommand(appsUpdateCmd)
	flags.UseNewImageFlag(appsUpdateCmd)
}

func updateApp(appName string, image *string) {
	client := config.Client()

	fmt.Printf("Pulling latest image for %s...\n", appName)
	_, err := client.Operations.UpdateApp(
		operations.NewUpdateAppParams().
			WithAppName(appName).
			WithNewImage(image),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("App reloaded!")
}
