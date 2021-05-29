package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func UseUpdateRunConfigFlags(cmd *cobra.Command) {
	cmd.Flags().Int64Slice("add-target-ports", []int64{}, "Add to the list of ports that the app is listening to inside the container")
	cmd.Flags().Int64Slice("rm-target-ports", []int64{}, "Remove from the list of ports that the app is listening to inside the container")

	cmd.Flags().Int64Slice("add-published-ports", []int64{}, "Add to the list of ports the app can be accessed through")
	cmd.Flags().Int64Slice("rm-published-ports", []int64{}, "Remove from the list of ports the app can be accessed through")

	cmd.Flags().StringSlice("add-placement-constraint", []string{}, "Add to the list of constraints specifying which nodes can run the app")
	cmd.Flags().StringSlice("rm-placement-constraint", []string{}, "Remove from the list of constraints specifying which nodes can run the app")
}

type UpdateRunConfig struct {
	AddTargetPorts         []uint32
	RMTargetPorts          []uint32
	AddPublishedPorts      []uint32
	RMPublishedPorts       []uint32
	AddPlacementConstraint []string
	RMPlacementConstraint  []string
}

func GetUpdateRunConfigFlags(cmd *cobra.Command) *UpdateRunConfig {
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

	return &UpdateRunConfig{
		AddTargetPorts:         addTargetPortsUInt32,
		RMTargetPorts:          rmTargetPortsUInt32,
		AddPublishedPorts:      addPublishedPortsUInt32,
		RMPublishedPorts:       rmPublishedPortsUInt32,
		AddPlacementConstraint: addPlacementConstraint,
		RMPlacementConstraint:  rmPlacementConstraint,
	}
}
