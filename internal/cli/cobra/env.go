package cobra

import (
	"context"
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/internal/utils"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print all the environment variables for an application",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		printEnv(appName)
	},
}

func init() {
	RootCmd.AddCommand(envCmd)
	flags.UseAppFlag(envCmd)
}

func printEnv(appName string) {
	ctx := context.Background()

	title.Printf("\nGetting environment for %s...\n", appName)
	app, err := api.GetApp(ctx, appName, `{ name, env }`)
	checkErr(err)
	env, err := getEnvEntries(ctx, app)
	checkErr(err)
	fmt.Println()

	lines := []string{}
	if len(env) == 0 {
		lines = append(lines, "<none>")
	} else {
		dimEnv := lo.Map(env, func(envVar string, _ int) string {
			return dim(envVar)
		})
		lines = append(lines, dimEnv...)
	}
	lines = append(lines, "")
	fmt.Println(strings.Join(lines, "\n"))
}

func getEnvEntries(ctx context.Context, app internal.App) ([]string, error) {

	return lo.Map(lo.Entries(utils.ToEnvMap(app.Env)), func(entry lo.Entry[string, string], _ int) string {
		return fmt.Sprintf("%s=%s", entry.Key, entry.Value)
	}), nil
}
