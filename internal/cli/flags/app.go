package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func UseAppFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("app", "a", "", "The app to perform the action on")
}

func GetAppFlag(cmd *cobra.Command) string {
	appName, err := cmd.Flags().GetString("app")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	appName = strings.TrimSpace(appName)
	if appName == "" {
		fmt.Println("Please pass a app name using the -a|--app flag")
		os.Exit(1)
	}

	return appName
}
