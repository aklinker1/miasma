package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseConfigFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("hidden", false, "Update the app to be hidden")
	cmd.Flags().Bool("rm-hidden", false, "Remove the app's hidden status")
	cmd.Flags().StringP("image", "i", "", "Change the image that the app runs")

	cmd.Flags().Int64Slice("add-target-ports", []int64{}, "Add to the list of ports that the app is listening to inside the container")
	cmd.Flags().Int64Slice("rm-target-ports", []int64{}, "Remove from the list of ports that the app is listening to inside the container")

	cmd.Flags().Int64Slice("add-published-ports", []int64{}, "Add to the list of ports the app can be accessed through")
	cmd.Flags().Int64Slice("rm-published-ports", []int64{}, "Remove from the list of ports the app can be accessed through")

	cmd.Flags().StringSlice("add-placement-constraint", []string{}, "Add to the list of constraints specifying which nodes can run the app")
	cmd.Flags().StringSlice("rm-placement-constraint", []string{}, "Remove from the list of constraints specifying which nodes can run the app")
}

type AppUpdateConfig struct {
	Hidden                 bool
	RMHidden               bool
	Image                  *string
	AddTargetPorts         []int64
	RMTargetPorts          []int64
	AddPublishedPorts      []int64
	RMPublishedPorts       []int64
	AddPlacementConstraint []string
	RMPlacementConstraint  []string
}

func GetConfigFlags(cmd *cobra.Command) *AppUpdateConfig {
	hidden, err := cmd.Flags().GetBool("hidden")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmHidden, err := cmd.Flags().GetBool("rm-hidden")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addTargetPorts, err := cmd.Flags().GetInt64Slice("add-target-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmTargetPorts, err := cmd.Flags().GetInt64Slice("rm-target-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addPublishedPorts, err := cmd.Flags().GetInt64Slice("add-published-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmPublishedPorts, err := cmd.Flags().GetInt64Slice("rm-published-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addPlacementConstraint, err := cmd.Flags().GetStringSlice("add-placement-constraint")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmPlacementConstraint, err := cmd.Flags().GetStringSlice("rm-placement-constraint")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var imagePtr *string
	if image != "" {
		imagePtr = &image
	}

	return &AppUpdateConfig{
		Hidden:                 hidden,
		RMHidden:               rmHidden,
		Image:                  imagePtr,
		AddTargetPorts:         addTargetPorts,
		RMTargetPorts:          rmTargetPorts,
		AddPublishedPorts:      addPublishedPorts,
		RMPublishedPorts:       rmPublishedPorts,
		AddPlacementConstraint: addPlacementConstraint,
		RMPlacementConstraint:  rmPlacementConstraint,
	}
}
