package cobra

import (
	"context"
	"strings"

	"github.com/aklinker1/miasma/internal"
	"github.com/aklinker1/miasma/internal/cli/flags"
	"github.com/spf13/cobra"
)

var appsEditCmd = &cobra.Command{
	Use:   "apps:edit",
	Short: "Update an app's display properties",
	Long: `Update an app's properties such as name, group, and target ports:

  miasma apps:edit --app app-name --group some-group

See the list of flags for all the properties that can be set for an application.
  
It is worth noting that for properties that are lists, there is no add or remove. Instead, include all the values for an array property you would like to change:
  
  miasma apps:configure --app app-name --ports 80,22
	  
Only the properties specified in the flags will update be updated. To remove a property, pass in an empty string for the value:
  
  miasma apps:configure --app app-name --ports ""
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		appName, deferable := flags.GetAppFlag(cmd)
		defer deferable()
		newConfig := flags.GetEditAppFlags(cmd)

		editApp(appName, newConfig)
	},
}

func init() {
	RootCmd.AddCommand(appsEditCmd)
	flags.UseAppFlag(appsEditCmd)
	flags.UseEditAppFlags(appsEditCmd)
}

func editApp(appName string, newApp *flags.EditApp) {
	ctx := context.Background()
	title.Printf("\nEditing %s...\n", appName)

	// Get current app
	app, err := api.GetApp(ctx, appName, `{
		id
		name
		command
		networks
		placement
		publishedPorts
		targetPorts
		volumes {
			target
			source
		}
	}`)
	checkErr(err)
	typedChanges := app

	// Merge lists
	existingTargetPortMap := map[int]int{}
	for index, port := range typedChanges.TargetPorts {
		existingTargetPortMap[port] = index
	}
	existingPublishedPortMap := map[int]int{}
	for index, port := range typedChanges.PublishedPorts {
		existingPublishedPortMap[port] = index
	}
	existingPlacementMap := map[string]int{}
	for index, placement := range typedChanges.Placement {
		existingPlacementMap[placement] = index
	}
	existingVolumesMap := map[string]int{}
	for index, volume := range typedChanges.Volumes {
		existingVolumesMap[volume.Source+":"+volume.Target] = index
	}

	// Update config
	for _, targetPort := range newApp.AddTargetPorts {
		if _, ok := existingTargetPortMap[targetPort]; !ok {
			typedChanges.TargetPorts = append(typedChanges.TargetPorts, targetPort)
			existingTargetPortMap[targetPort] = len(typedChanges.TargetPorts) - 1
		}
	}
	for _, targetPort := range newApp.RMTargetPorts {
		if index, ok := existingTargetPortMap[targetPort]; ok {
			typedChanges.TargetPorts = append(
				typedChanges.TargetPorts[:index],
				typedChanges.TargetPorts[index+1:]...,
			)
		}
	}
	for _, publishedPort := range newApp.AddPublishedPorts {
		if _, ok := existingPublishedPortMap[publishedPort]; !ok {
			typedChanges.PublishedPorts = append(typedChanges.PublishedPorts, publishedPort)
			existingPublishedPortMap[publishedPort] = len(typedChanges.PublishedPorts) - 1
		}
	}
	for _, publishedPort := range newApp.RMPublishedPorts {
		if index, ok := existingPublishedPortMap[publishedPort]; ok {
			typedChanges.PublishedPorts = append(
				typedChanges.PublishedPorts[:index],
				typedChanges.PublishedPorts[index+1:]...,
			)
		}
	}
	for _, placement := range newApp.AddPlacementConstraint {
		if _, ok := existingPlacementMap[placement]; !ok {
			typedChanges.Placement = append(typedChanges.Placement, placement)
			existingPlacementMap[placement] = len(typedChanges.Placement) - 1
		}
	}
	for _, placement := range newApp.RMPlacementConstraint {
		if index, ok := existingPlacementMap[placement]; ok {
			typedChanges.Placement = append(
				typedChanges.Placement[:index],
				typedChanges.Placement[index+1:]...,
			)
		}
	}
	for _, volume := range newApp.AddVolume {
		if _, ok := existingVolumesMap[volume]; !ok {
			split := strings.SplitN(volume, ":", 2)
			source := split[0]
			target := split[1]
			typedChanges.Volumes = append(typedChanges.Volumes, &internal.BoundVolume{
				Target: target,
				Source: source,
			})
			existingVolumesMap[volume] = len(typedChanges.Volumes) - 1
		}
	}
	for _, volume := range newApp.RMVolume {
		if index, ok := existingVolumesMap[volume]; ok {
			typedChanges.Volumes = append(
				typedChanges.Volumes[:index],
				typedChanges.Volumes[index+1:]...,
			)
		}
	}

	changes := map[string]any{
		"command":        typedChanges.Command,
		"networks":       typedChanges.Networks,
		"placement":      typedChanges.Placement,
		"publishedPorts": typedChanges.PublishedPorts,
		"targetPorts":    typedChanges.TargetPorts,
		"volumes":        typedChanges.Volumes,
	}

	// Simple updates
	if newApp.Hidden {
		changes["hidden"] = true
	} else if newApp.Visible {
		changes["hidden"] = false
	}
	if newApp.Name != nil {
		changes["name"] = *newApp.Name
	}
	if newApp.Group != nil {
		changes["group"] = *newApp.Group
	}

	// push app updates
	_, err = api.EditApp(ctx, typedChanges.ID, changes, `{ updatedAt }`)
	checkErr(err)

	done("%s updated", appName)
}
