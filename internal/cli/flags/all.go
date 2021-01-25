package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseAllFlag(cmd *cobra.Command) {
	cmd.Flags().BoolP("all", "A", false, "List all apps, including hidden ones")
}

func GetAllFlag(cmd *cobra.Command) bool {
	all, err := cmd.Flags().GetBool("all")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return all
}
