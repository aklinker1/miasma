package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var shouldShowVersion bool

func UseVersionFlag(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&shouldShowVersion, "version", "v", false, "Print the CLI version")
}

func GetVersionFlag(cmd *cobra.Command) bool {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return version
}
