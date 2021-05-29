package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aklinker1/miasma/package/models"
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

func AppInput(app *models.AppInput) error {
	err := AppName(app.Name)
	if err != nil {
		return nil
	}
	if len(app.Image) == 0 {
		return errors.New("'image' is required")
	}
	return nil
}
