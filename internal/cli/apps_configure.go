package cli

import (
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsConfigureCmd = &cobra.Command{
	Use:   "apps:configure",
	Short: "Update an application's properties",
	Long: `Update an application's properties such as target ports, networks, etc. See the list of
flags for all the properties that can be set for an application.

It is worth noting that for properties that are lists, there is no add or remove. Instead, include
all the values for an array property you would like to change:

  miasma app:configure --app app-name --ports 80,22
	
Only the properties specified in the flags will update be updated. To remove a propterty, pass in an empty string for the value:

  miasma app:configure --app app-name --ports ""`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		stopApp(appName)
	},
}

func init() {
	RootCmd.AddCommand(appsConfigureCmd)
	flags.UseAppFlag(appsConfigureCmd)
}

func configureApp(appName string) {
	panic("NOT IMPLEMENTED")
	// Update
}
