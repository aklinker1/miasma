package app_service

import (
	"github.com/aklinker1/miasma/internal/server/services/docker_service"
	"github.com/aklinker1/miasma/internal/server/utils/server_models"
	"github.com/aklinker1/miasma/internal/shared/log"
)

// Upgrade pulls the latest image and reloads the application
func Upgrade(
	details *server_models.AppDetails,
	env map[string]string,
	plugins *server_models.AppPlugins,
	newImage string,
) (bool, error) {
	log.V("app_service.Start(%v, env:***, %v, %v)", details, plugins, newImage)
	originalImage := details.App.Image
	originalDigest := details.RunConfig.ImageDigest
	log.V("Image: '%s' -> '%s'", originalImage, newImage)

	// Pull and get the new digest
	err := docker_service.PullImage(newImage)
	if err != nil {
		return false, err
	}
	newDigest, err := docker_service.GetImageDigest(newImage)
	if err != nil {
		return false, err
	}
	log.V("Digest: '%s' -> '%s'", originalDigest, newDigest)

	// Return if we aren't actually updating
	details.App.Image = newImage
	details.RunConfig.ImageDigest = newDigest
	if newDigest == originalDigest {
		return false, nil
	}

	err = Reload(details, env, plugins)
	if err != nil {
		return false, err
	}

	// Update the app's image if the image changed
	if originalImage != newImage {
		panic("TODO: Save app and run config")
		// err = app_service.WriteAppMeta(appMeta)
		// if err != nil {
		// 	log.W("Failed to save new app image after updating the app")
		// }
	}

	return true, nil
}
