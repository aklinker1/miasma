package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func AppName(appName string) error {
	appName = strings.TrimSpace(appName)
	if len(appName) < 3 {
		return errors.New("App name must be at least 3 characters long")
	}
	validCharsStr := "^[a-z][a-z0-9.-]+[a-z0-9]$"
	validChars := regexp.MustCompile(validCharsStr)
	if validChars.FindString(appName) == "" {
		return fmt.Errorf("'%s' must match '%s'", appName, validCharsStr)
	}

	return nil
}
