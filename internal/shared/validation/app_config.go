package validation

import (
	"errors"

	"github.com/aklinker1/miasma/internal/server/gen/models"
)

func AppConfig(appConfig *models.AppConfig) error {
	// Networks
	// Placement
	// PublishedPorts
	// TargetPorts

	// Route
	if appConfig.Route != nil {
		if appConfig.Route.TraefikRule != nil && appConfig.Route.Host != nil {
			return errors.New("Route cannot include both a 'traefikRule' and 'host'. Only one can be defined")
		} else if appConfig.Route.Host == nil && appConfig.Route.Path != nil {
			return errors.New("Route can only include a 'path' when the 'host' is also provided")
		}
	}

	return nil
}
