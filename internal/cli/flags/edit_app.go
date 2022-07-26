package flags

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func UseEditAppFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("name", "n", "", "Change the app's name")

	cmd.Flags().StringP("group", "g", "", "Change the app's group")

	cmd.Flags().Bool("hidden", false, "Make the app hidden")
	cmd.Flags().Bool("visible", false, "Remove the hidden flag from the app")

	cmd.Flags().IntSlice("add-target-ports", []int{}, "Add to the list of ports that the app is listening to inside the container")
	cmd.Flags().IntSlice("rm-target-ports", []int{}, "Remove from the list of ports that the app is listening to inside the container")

	cmd.Flags().IntSlice("add-published-ports", []int{}, "Add to the list of ports the app can be accessed through")
	cmd.Flags().IntSlice("rm-published-ports", []int{}, "Remove from the list of ports the app can be accessed through")

	cmd.Flags().StringSlice("add-placement-constraint", []string{}, "Add to the list of constraints specifying which nodes can run the app")
	cmd.Flags().StringSlice("rm-placement-constraint", []string{}, "Remove from the list of constraints specifying which nodes can run the app")

	cmd.Flags().StringSlice("add-volume", []string{}, "Add a bound volume to the host machine")
	cmd.Flags().StringSlice("rm-volume", []string{}, "Remove a bound volume to the host machine")
}

type EditApp struct {
	Hidden                 bool
	Visible                bool
	Name                   *string
	Group                  *string
	AddTargetPorts         []int
	RMTargetPorts          []int
	AddPublishedPorts      []int
	RMPublishedPorts       []int
	AddPlacementConstraint []string
	RMPlacementConstraint  []string
	AddVolume              []string
	RMVolume               []string
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

	// Target Ports
	addTargetPorts, err := cmd.Flags().GetIntSlice("add-target-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmTargetPorts, err := cmd.Flags().GetIntSlice("rm-target-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Published Ports
	addPublishedPorts, err := cmd.Flags().GetIntSlice("add-published-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmPublishedPorts, err := cmd.Flags().GetIntSlice("rm-published-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Placement
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

	// Volumes
	addVolume, err := cmd.Flags().GetStringSlice("add-volume")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmVolume, err := cmd.Flags().GetStringSlice("rm-volume")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &EditApp{
		Hidden:                 hidden,
		Visible:                visible,
		Name:                   namePtr,
		Group:                  groupPtr,
		AddTargetPorts:         addTargetPorts,
		RMTargetPorts:          rmTargetPorts,
		AddPublishedPorts:      addPublishedPorts,
		RMPublishedPorts:       rmPublishedPorts,
		AddPlacementConstraint: addPlacementConstraint,
		RMPlacementConstraint:  rmPlacementConstraint,
		AddVolume:              addVolume,
		RMVolume:               rmVolume,
	}
}
