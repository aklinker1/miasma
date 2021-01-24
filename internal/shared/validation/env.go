package validation

import (
	"fmt"
	"regexp"
)

func AppEnv(env map[string]interface{}) error {
	portRegex := regexp.MustCompile("^PORT(_[1-9][0-9]*)?$")
	for key, _ := range env {
		if portRegex.FindString(key) != "" {
			return fmt.Errorf("%s is not allowed, env vars can't be named PORT or PORT_XX.\nThese variables are already provided based on the app's target adn published ports", key)
		}
	}

	return nil
}
