package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseHiddenFlag(cmd *cobra.Command) {
	cmd.Flags().Bool("hidden", false, "Whether or not the app is hidden")
}

func GetHiddenFlag(cmd *cobra.Command) bool {
	hidden, err := cmd.Flags().GetBool("hidden")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return hidden
}
