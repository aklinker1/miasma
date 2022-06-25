package cobra

import (
	"context"
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal/cli"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/fatih/color"
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

func listApps(showHidden bool) {
	ctx := context.Background()
	apps, err := api.ListApps(
		ctx,
		cli.ListAppsOptions{
			ShowHidden: &showHidden,
		},
		`{
			name
			group
			status
			instances {
				running
				total
			}
			simpleRoute
		}`,
	)
	checkErr(err)

	title.Println("\nApps")
	prevGroup := ""
	for _, app := range apps {
		if app.Group != nil && *app.Group != prevGroup {
			title.Printf("\n%s\n", *app.Group)
		}
		status := green("●")
		if app.Status == "stopped" {
			status = red("●")
		}
		if app.Instances != nil && app.Instances.Total > 0 && app.Instances.Running != app.Instances.Total {
			status = yellow("●")
		}
		appDetails := []string{}
		if app.Instances != nil {
			appDetails = append(
				appDetails,
				fmt.Sprintf("%d/%d running", app.Instances.Running, app.Instances.Total),
			)
		}
		if app.SimpleRoute != nil {
			appDetails = append(appDetails, color.New(color.Underline).Sprint(*app.SimpleRoute))
		}
		appDetailsStr := fmt.Sprintf("(%s)", strings.Join(appDetails, ", "))
		fmt.Printf(" %s %s %s\n", status, bold(app.Name), dim(appDetailsStr))
		if app.Group == nil {
			prevGroup = ""
		} else {
			prevGroup = *app.Group
		}
	}
	fmt.Printf(dim("\n(%d total)\n\n"), len(apps))
}
