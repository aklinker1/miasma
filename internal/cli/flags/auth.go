package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseAuthFlag(cmd *cobra.Command) {
	cmd.Flags().String("auth", "", "The server's access token if one is setup")
}

func GetAuthFlag(cmd *cobra.Command) string {
	v, err := cmd.Flags().GetString("auth")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return v
}
