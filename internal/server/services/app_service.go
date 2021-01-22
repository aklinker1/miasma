package services

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/mappers"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared/log"

	"gopkg.in/yaml.v2"
)

type appService struct{}

var App = &appService{}

// App

func (service *appService) AppsDir() (dir string, err error) {
	dir = "/data/miasma/apps"
	if exists, _ := Files.dirExists(dir); !exists {
		// https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value
		err = os.MkdirAll(dir, 0755)
	}
	return dir, err
}

func (service *appService) Get(appName string) (*models.App, error) {
	appsDir, err := service.AppsDir()
	if err != nil {
		return nil, err
	}
	appDir := fmt.Sprintf("%s/%s", appsDir, appName)
	if exists, _ := Files.dirExists(appDir); !exists {
		return nil, fmt.Errorf("%s is not an application. Have you ran `miasma app:create %s`?", appName, appName)
	}
	metaFilePath := fmt.Sprintf("%s/%s", appDir, "meta.yml")

	metaFile, err := ioutil.ReadFile(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("Could not find data for %s, did %s get moved?", appName, metaFilePath)
	}

	var metaYml = &types.AppMetaData{}
	if err := yaml.Unmarshal(metaFile, metaYml); err != nil {
		return nil, err
	}

	return mappers.App.FromMeta(appName, metaYml, Docker.IsAppServiceRunning(appName)), err
}

func (service *appService) GetAll(showHidden bool) ([]*models.App, error) {
	appsDir, err := service.AppsDir()
	if err != nil {
		return nil, err
	}
	appDirs, err := ioutil.ReadDir(appsDir)
	if err != nil {
		return nil, err
	}

	result := []*models.App{}
	for _, appDir := range appDirs {
		app, err := service.Get(appDir.Name())
		if err != nil {
			return nil, err
		}
		if showHidden || !app.Hidden {
			result = append(result, app)
		}
	}
	return result, nil
}

func (service *appService) Create(app models.AppInput) (*models.App, error) {
	appsDir, err := service.AppsDir()
	if err != nil {
		return nil, err
	}
	appDir := fmt.Sprintf("%s/%s", appsDir, *app.Name)
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		return nil, err
	}
	metaPath := fmt.Sprintf("%s/meta.yml", appDir)
	metaData, err := yaml.Marshal(mappers.App.ToMeta(&app))
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(metaPath, metaData, 0755)
	if err != nil {
		return nil, err
	}

	return service.Get(*app.Name)
}

func (service *appService) Delete(app *models.App) error {
	appsDir, err := service.AppsDir()
	if err != nil {
		return err
	}

	appDir := fmt.Sprintf("%s/%s", appsDir, *app.Name)
	err = os.RemoveAll(appDir)
	if err != nil {
		return err
	}

	return nil
}
