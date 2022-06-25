package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseNewImageFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("image", "i", "", "The image the app should use instead of the current image")
}

func GetNewImageFlag(cmd *cobra.Command) *string {
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if image == "" {
		return nil
	}
	return &image
}
