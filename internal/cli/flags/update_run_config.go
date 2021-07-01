package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func UseEditAppFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("hidden", false, "Make the app hidden")
	cmd.Flags().Bool("visible", false, "Remove the hidden flag from the app")
	cmd.Flags().StringP("name", "n", "", "Change the app's name")
	cmd.Flags().StringP("group", "g", "", "Change the app's group")
}

type EditApp struct {
	Hidden  bool
	Visible bool
	Name    *string
	Group   *string
}

func GetEditAppFlags(cmd *cobra.Command) *EditApp {
	// Group
	group, err := cmd.Flags().GetString("group")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	groupTrimmed := strings.TrimSpace(group)
	var groupPtr *string
	if groupTrimmed != "" {
		groupPtr = &groupTrimmed
	}

	// Name
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	nameTrimmed := strings.TrimSpace(name)
	var namePtr *string
	if nameTrimmed != "" {
		namePtr = &nameTrimmed
	}

	// Hidden
	hidden, err := cmd.Flags().GetBool("hidden")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	visible, err := cmd.Flags().GetBool("visible")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &EditApp{
		Hidden:  hidden,
		Visible: visible,
		Name:    namePtr,
		Group:   groupPtr,
	}
}
