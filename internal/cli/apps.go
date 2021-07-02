package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var purple = color.New(color.FgMagenta).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()
var dim = color.New(color.Faint).SprintFunc()

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
	client := config.Client()
	apps, err := client.Operations.ListApps(operations.NewListAppsParams().WithHidden(&includeHidden))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	color.Magenta("Apps")
	prevGroup := ""
	for _, app := range apps.Payload {
		if app.Group != prevGroup && app.Group != "" {
			fmt.Println()
			color.Magenta(app.Group)
		}
		appDetails := "" // "(:3000, 0/1)"
		fmt.Printf(" %s %s %s\n", green("●"), bold(app.Name), dim(appDetails))
		prevGroup = app.Group
	}
	fmt.Printf(dim("\n(%d total)\n"), len(apps.Payload))
}
