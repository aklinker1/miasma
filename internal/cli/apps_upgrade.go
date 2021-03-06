package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var appsUpgradeCmd = &cobra.Command{
	Use:   "apps:upgrade",
	Short: "Pull the latest version of the application's image and reload the app",
	Long:  "Pull the latest version of the application's image and reload the app. If a new image is passed, the app is updated to use that image instead of the current one",
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()
		image := flags.GetNewImageFlag(cmd)

		updateApp(appName, image)
	},
}

func init() {
	RootCmd.AddCommand(appsUpgradeCmd)
	flags.UseNewImageFlag(appsUpgradeCmd)
	flags.UseAppFlag(appsUpgradeCmd)
}

func updateApp(appName string, image *string) {
	client := config.Client()

	fmt.Printf("Pulling latest image for %s...\n", appName)
	_, err := client.Operations.UpgradeApp(
		operations.NewUpgradeAppParams().
			WithAppName(appName).
			WithNewImage(image),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("App reloaded!")
}
