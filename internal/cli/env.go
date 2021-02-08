package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
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
	fmt.Printf("Getting env for %s...\n", appName)
	responseEnv := getEnvText(appName)
	lines := []string{
		"Done!",
		"",
		"Environment Variables:",
	}
	if len(responseEnv) == 0 {
		lines = append(lines, "<none>")
	} else {
		lines = append(lines, responseEnv...)
	}
	lines = append(lines, "")
	fmt.Println(strings.Join(lines, "\n"))
}

func getEnvText(appName string) []string {
	getEnvParams := operations.NewGetAppEnvParams()
	getEnvParams.AppName = appName
	response, err := config.Client().Operations.GetAppEnv(getEnvParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := []string{}
	for key, value := range response.Payload.(map[string]interface{}) {
		lines = append(lines, fmt.Sprintf("%s=%v", key, value))
	}
	return lines
}
