package validation

import (
	"errors"

	"github.com/aklinker1/miasma/package/models"
)

func TraefikPluginConfig(traefik *models.InputTraefikPluginConfig) error {
	if traefik.TraefikRule != nil && traefik.Host != nil {
		return errors.New("Route cannot include both a 'traefikRule' and 'host'. Only one can be defined")
	} else if traefik.Host == nil && traefik.Path != nil {
		return errors.New("Route can only include a 'path' when the 'host' is also provided")
	}

	return nil
}
