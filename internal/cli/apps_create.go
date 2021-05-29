package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/internal/shared/validation"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/aklinker1/miasma/package/models"
	"github.com/spf13/cobra"
)

var appsCreateCmd = &cobra.Command{
	Use:   "apps:create",
	Short: "Create and deploy a new application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		err := validation.AppName(appName)
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
	client := config.Client()

	fmt.Printf("Creating %s...\n", appName)
	createParams := operations.NewCreateAppParams()
	createParams.App = &models.AppInput{
		Name:   appName,
		Image:  image,
		Hidden: hidden,
	}
	_, err := client.Operations.CreateApp(createParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Starting...")
	startParams := operations.NewStartAppParams()
	startParams.AppName = appName
	_, err = client.Operations.StartApp(startParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}
