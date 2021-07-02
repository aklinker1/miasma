package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseUpdateRunConfigFlags(cmd *cobra.Command) {
	cmd.Flags().Int32Slice("add-target-ports", []int32{}, "Add to the list of ports that the app is listening to inside the container")
	cmd.Flags().Int32Slice("rm-target-ports", []int32{}, "Remove from the list of ports that the app is listening to inside the container")

	cmd.Flags().Int32Slice("add-published-ports", []int32{}, "Add to the list of ports the app can be accessed through")
	cmd.Flags().Int32Slice("rm-published-ports", []int32{}, "Remove from the list of ports the app can be accessed through")

	cmd.Flags().StringSlice("add-placement-constraint", []string{}, "Add to the list of constraints specifying which nodes can run the app")
	cmd.Flags().StringSlice("rm-placement-constraint", []string{}, "Remove from the list of constraints specifying which nodes can run the app")

	cmd.Flags().StringSlice("add-volume", []string{}, "Add a bound volume to the host machine")
	cmd.Flags().StringSlice("rm-volume", []string{}, "Remove a bound volume to the host machine")
}

type ConfigureApp struct {
	AddTargetPorts         []uint32
	RMTargetPorts          []uint32
	AddPublishedPorts      []uint32
	RMPublishedPorts       []uint32
	AddPlacementConstraint []string
	RMPlacementConstraint  []string
	AddVolume              []string
	RMVolume               []string
}

func GetConfigureAppFlags(cmd *cobra.Command) *ConfigureApp {
	addTargetPorts, err := cmd.Flags().GetInt32Slice("add-target-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	addTargetPortsUInt32 := make([]uint32, len(addTargetPorts))
	for i, port := range addTargetPorts {
		addTargetPortsUInt32[i] = uint32(port)
	}
	rmTargetPorts, err := cmd.Flags().GetInt32Slice("rm-target-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmTargetPortsUInt32 := make([]uint32, len(rmTargetPorts))
	for i, port := range rmTargetPorts {
		rmTargetPortsUInt32[i] = uint32(port)
	}

	addPublishedPorts, err := cmd.Flags().GetInt32Slice("add-published-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	addPublishedPortsUInt32 := make([]uint32, len(addPublishedPorts))
	for i, port := range addPublishedPorts {
		addPublishedPortsUInt32[i] = uint32(port)
	}
	rmPublishedPorts, err := cmd.Flags().GetInt32Slice("rm-published-ports")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rmPublishedPortsUInt32 := make([]uint32, len(rmPublishedPorts))
	for i, port := range rmPublishedPorts {
		rmPublishedPortsUInt32[i] = uint32(port)
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

	return &ConfigureApp{
		AddTargetPorts:         addTargetPortsUInt32,
		RMTargetPorts:          rmTargetPortsUInt32,
		AddPublishedPorts:      addPublishedPortsUInt32,
		RMPublishedPorts:       rmPublishedPortsUInt32,
		AddPlacementConstraint: addPlacementConstraint,
		RMPlacementConstraint:  rmPlacementConstraint,
		AddVolume:              addVolume,
		RMVolume:               rmVolume,
	}
}
