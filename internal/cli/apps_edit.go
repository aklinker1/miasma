package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/aklinker1/miasma/package/models"
	"github.com/spf13/cobra"
)

var appsEditCmd = &cobra.Command{
	Use:   "apps:edit",
	Short: "Update an app's display properties",
	Long: `Update an app's properties such as name and group. See the list of flags for all the properties
that can be set for an application.

  miasma apps:edit --app app-name --group some-group
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()
		newConfig := flags.GetEditAppFlags(cmd)

		editApp(appName, newConfig)
	},
}

func init() {
	RootCmd.AddCommand(appsEditCmd)
	flags.UseAppFlag(appsEditCmd)
	flags.UseEditAppFlags(appsEditCmd)
}

func editApp(appName string, newApp *flags.EditApp) {
	fmt.Printf("Editing %s...\n", appName)

	// Get current app
	client := config.Client()
	editResponse, err := client.Operations.GetApp(
		operations.NewGetAppParams().WithAppName(appName),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app := &models.AppEdit{
		Name:   editResponse.Payload.Name,
		Group:  editResponse.Payload.Group,
		Hidden: editResponse.Payload.Hidden,
	}

	// Update existing app
	if newApp.Hidden {
		app.Hidden = true
	} else if newApp.Visible {
		app.Hidden = false
	}
	if newApp.Name != nil {
		app.Name = *newApp.Name
	}
	if newApp.Group != nil {
		app.Group = *newApp.Group
	}

	// push app updates
	_, err = client.Operations.EditApp(
		operations.NewEditAppParams().
			WithAppName(appName).
			WithNewApp(app),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
