package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/editor"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/spf13/cobra"
)

var envEditCmd = &cobra.Command{
	Use:   "env:edit",
	Short: "Edit the environment variables for the application",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()

		editEnv(appName)
	},
}

func init() {
	RootCmd.AddCommand(envEditCmd)
	flags.UseAppFlag(envEditCmd)
}

func editEnv(appName string) {
	fmt.Printf("Getting env for %s...\n", appName)
	lines := []string{"# Do not escape any characters or add quotes around the value"}
	lines = append(lines, getEnvText(appName)...)
	fmt.Println("Done!")
	newText, err := editor.EditText(strings.Join(lines, "\n"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newLines := strings.Split(newText, "\n")
	newEnv := map[string]string{}
	for lineIndex, line := range newLines {
		trimmedLine := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmedLine, "#") && trimmedLine != "" {
			keyValue := strings.SplitN(trimmedLine, "=", 2)
			fmt.Println(1, trimmedLine, 2, keyValue)
			if len(keyValue) != 2 {
				fmt.Println(strings.Join([]string{
					"Invalid environment variable:",
					"",
					fmt.Sprintf("  Line %d: %s", lineIndex+1, trimmedLine),
					"",
					"Make sure it is formatted as KEY=VALUE",
				}, "\n"))
				os.Exit(1)
			}
			newEnv[keyValue[0]] = keyValue[1]
		}
	}
	fmt.Printf("Updating env for %s...\n", appName)
	params := operations.NewUpdateAppEnvParams().
		WithAppName(appName).
		WithNewEnv(newEnv)
	_, err = config.Client().Operations.UpdateAppEnv(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
