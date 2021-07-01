package cli

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/cli/config"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/aklinker1/miasma/package/client/operations"
	"github.com/aklinker1/miasma/package/models"
	"github.com/spf13/cobra"
)

var appsConfigureCmd = &cobra.Command{
	Use:   "apps:configure",
	Short: "Update an application's properties",
	Long: `Update an application's properties such as target ports. See the list of flags for all the properties
that can be set for an application.

It is worth noting that for properties that are lists, there is no add or remove. Instead, include
all the values for an array property you would like to change:

  miasma apps:configure --app app-name --ports 80,22
	
Only the properties specified in the flags will update be updated. To remove a propterty, pass in an empty string for the value:

  miasma apps:configure --app app-name --ports ""`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()
		newConfig := flags.GetUpdateRunConfigFlags(cmd)

		configureApp(appName, newConfig)
	},
}

func init() {
	RootCmd.AddCommand(appsConfigureCmd)
	flags.UseAppFlag(appsConfigureCmd)
	flags.UseUpdateRunConfigFlags(appsConfigureCmd)
}

func configureApp(appName string, newConfig *flags.UpdateRunConfig) {
	fmt.Printf("Updating %s...\n", appName)

	// Get current config
	client := config.Client()
	configResponse, err := client.Operations.GetRunConfig(
		operations.NewGetRunConfigParams().WithAppName(appName),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config := &models.InputRunConfig{
		Command:        configResponse.Payload.Command,
		Networks:       configResponse.Payload.Networks,
		Placement:      configResponse.Payload.Placement,
		PublishedPorts: configResponse.Payload.PublishedPorts,
		TargetPorts:    configResponse.Payload.TargetPorts,
		Volumes:        configResponse.Payload.Volumes,
	}
	existingTargetPortMap := map[uint32]int{}
	for index, port := range config.TargetPorts {
		existingTargetPortMap[port] = index
	}
	existingPublishedPortMap := map[uint32]int{}
	for index, port := range config.PublishedPorts {
		existingPublishedPortMap[port] = index
	}
	existingPlacementMap := map[string]int{}
	for index, placement := range config.Placement {
		existingPlacementMap[placement] = index
	}

	// Update config
	// if newConfig.Hidden {
	// 	config.Hidden = true
	// } else if newConfig.RMHidden {
	// 	config.Hidden = false
	// }
	// if newConfig.Image != nil {
	// 	config.Image = *newConfig.Image
	// }
	for _, targetPort := range newConfig.AddTargetPorts {
		if _, ok := existingTargetPortMap[targetPort]; !ok {
			config.TargetPorts = append(config.TargetPorts, targetPort)
			existingTargetPortMap[targetPort] = len(config.TargetPorts) - 1
		}
	}
	for _, targetPort := range newConfig.RMTargetPorts {
		if index, ok := existingTargetPortMap[targetPort]; ok {
			config.TargetPorts = append(
				config.TargetPorts[:index],
				config.TargetPorts[index+1:]...,
			)
		}
	}
	for _, publishedPort := range newConfig.AddPublishedPorts {
		if _, ok := existingPublishedPortMap[publishedPort]; !ok {
			config.PublishedPorts = append(config.PublishedPorts, publishedPort)
			existingPublishedPortMap[publishedPort] = len(config.PublishedPorts) - 1
		}
	}
	for _, publishedPort := range newConfig.RMPublishedPorts {
		if index, ok := existingPublishedPortMap[publishedPort]; ok {
			config.PublishedPorts = append(
				config.PublishedPorts[:index],
				config.PublishedPorts[index+1:]...,
			)
		}
	}
	for _, placement := range newConfig.AddPlacementConstraint {
		if _, ok := existingPlacementMap[placement]; !ok {
			config.Placement = append(config.Placement, placement)
			existingPlacementMap[placement] = len(config.Placement) - 1
		}
	}
	for _, placement := range newConfig.RMPlacementConstraint {
		if index, ok := existingPlacementMap[placement]; ok {
			config.Placement = append(
				config.Placement[:index],
				config.Placement[index+1:]...,
			)
		}
	}

	// push config updates
	_, err = client.Operations.UpdateRunConfig(
		operations.NewUpdateRunConfigParams().
			WithAppName(appName).
			WithNewRunConfig(config),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done!")
}
