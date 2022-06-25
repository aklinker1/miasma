package cobra

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/cli/editor"
	"github.com/aklinker1/miasma/internal/cli/flags"
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
	ctx := context.Background()
	title.Printf("\nSetting %s's env...\n", appName)

	app, err := api.GetApp(ctx, appName, `{ id, name, env }`)
	checkErr(err)
	env, err := getEnvEntries(ctx, app)
	checkErr(err)

	// Edit the env
	fmt.Println("Edit env...")
	initialLines := []string{
		"# Do not escape any characters or add quotes around the value",
	}
	initialLines = append(initialLines, env...)
	initialText := strings.Join(initialLines, "\n")
	newEnv, err := editEnvInEditor(ctx, initialText)
	checkErr(err)

	// Update app
	fmt.Println("Updating app...")
	err = api.SetAppEnv(ctx, app.ID, newEnv)
	checkErr(err)

	fmt.Printf("Done!\n\n")
}

func editEnvInEditor(ctx context.Context, text string) (internal.EnvMap, error) {
	newText, err := editor.EditText(text)
	if err != nil {
		return nil, err
	}

	newLines := strings.Split(newText, "\n")
	newEnv := internal.EnvMap{}
	for lineIndex, line := range newLines {
		trimmedLine := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmedLine, "#") && trimmedLine != "" {
			keyValue := strings.SplitN(trimmedLine, "=", 2)
			if len(keyValue) != 2 {
				return nil, errors.New(strings.Join([]string{
					"Invalid environment variable:",
					"",
					fmt.Sprintf("  Line %d: %s", lineIndex+1, trimmedLine),
					"",
					"Make sure it is formatted as KEY=VALUE",
				}, "\n"))
			} else {
				newEnv[keyValue[0]] = keyValue[1]
			}
		}
	}
	return newEnv, nil
}
