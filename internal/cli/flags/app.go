package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func UseAppFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("app", "a", "", "The app to perform the action on")
}

func GetAppFlag(cmd *cobra.Command) (appName string, deferable func()) {
	appName, err := cmd.Flags().GetString("app")
	if err != nil || appName == "" {
		appName = viper.GetString("app")
		if appName == "" {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	appName = strings.TrimSpace(appName)
	if appName == "" {
		fmt.Println("Please pass a app name using the -a|--app flag")
		os.Exit(1)
	}

	return appName, func() {
		viper.Set("app", appName)
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
}
