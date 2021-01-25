package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseImageFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("image", "i", "", "Application's docker image that is ran")
}

func GetImageFlag(cmd *cobra.Command) string {
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return image
}
