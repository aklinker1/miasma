package cobra

import (
	"context"
	"fmt"

	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsUpgradeCmd = &cobra.Command{
	Use:   "apps:upgrade",
	Short: "Pull the latest version of the application's image and restart the app",
	Long:  "Pull the latest version of the application's image and restart the app. To change the image to a different tag or different image altogether, use apps:edit instead",
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		upgradeApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsUpgradeCmd)
	flags.UseAppFlag(appsUpgradeCmd)
}

func upgradeApp(appName string) {
	ctx := context.Background()
	title.Printf("\nUpgrading %s...\n", appName)

	oldApp, err := api.GetApp(ctx, appName, `{
		id
		name
		image
		imageDigest
	}`)
	checkErr(err)

	fmt.Printf("Pulling latest image for %s...\n", oldApp.Image)
	newApp, err := api.EditApp(
		ctx,
		oldApp.ID,
		map[string]any{
			"image": oldApp.Image,
		},
		`{ imageDigest }`,
	)
	checkErr(err)

	fmt.Printf(dim("Old digest: %s\n"), oldApp.ImageDigest)
	fmt.Printf(dim("New digest: %s\n"), newApp.ImageDigest)
	if newApp.ImageDigest != oldApp.ImageDigest {
		done("%s upgraded", appName)
	} else {
		done("%s is already up to date", appName)
	}

	// TODO show message for when upgrade didn't do anything
}
