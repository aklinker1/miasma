package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aklinker1/miasma/internal/server/gen/models"
	"github.com/aklinker1/miasma/internal/server/utils/mappers"
	"github.com/aklinker1/miasma/internal/server/utils/types"
	"github.com/aklinker1/miasma/internal/shared"
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

func (service *appService) GetAppMeta(appName string) (*types.AppMetaData, error) {
	appsDir, err := service.AppsDir()
	if err != nil {
		return nil, err
	}
	metaFilePath := fmt.Sprintf("%s/%s.yml", appsDir, appName)

	metaFile, err := ioutil.ReadFile(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("Could not find data for %s, did %s get moved?", appName, metaFilePath)
	}

	var metaYml = &types.AppMetaData{}
	if err := yaml.Unmarshal(metaFile, metaYml); err != nil {
		return nil, err
	}
	metaYml.Name = appName
	return metaYml, nil
}

func (service *appService) WriteAppMeta(appMeta *types.AppMetaData) error {
	appsDir, err := service.AppsDir()
	if err != nil {
		return err
	}
	metaFilePath := fmt.Sprintf("%s/%s.yml", appsDir, appMeta.Name)

	data, err := yaml.Marshal(appMeta)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(metaFilePath, data, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (service *appService) Get(appName string) (*models.App, error) {
	metaYml, err := service.GetAppMeta(appName)
	if err != nil {
		return nil, err
	}
	return mappers.App.FromMeta(appName, metaYml, Docker.IsAppServiceRunning(appName)), nil
}

func (service *appService) GetAll(showHidden bool) ([]*models.App, error) {
	appsDir, err := service.AppsDir()
	if err != nil {
		return nil, err
	}
	metaFiles, err := ioutil.ReadDir(appsDir)
	if err != nil {
		return nil, err
	}

	result := []*models.App{}
	for _, metaFile := range metaFiles {
		appName := strings.Replace(metaFile.Name(), ".yml", "", 1)
		log.V("%s > %s", metaFile.Name(), appName)
		app, err := service.Get(appName)
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
	metaPath := fmt.Sprintf("%s/%s.yml", appsDir, *app.Name)
	metaData, err := yaml.Marshal(mappers.App.ToMeta(&app))
	if err != nil {
		return nil, err
	}
	err = Docker.CreateNetwork(*app.Name)
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
	if err := Docker.DestroyNetwork(*app.Name); err != nil {
		log.W("Failed to destroy network: %v", err)
	}
	err := Docker.StopApp(app)
	if err != nil {
		log.V("%s is not running, no need to stop it (%v)", *app.Name, err)
	} else {
		log.V("Stopped %s before deleting", *app.Name)
	}

	appsDir, err := service.AppsDir()
	if err != nil {
		return err
	}

	metaPath := fmt.Sprintf("%s/%s.yml", appsDir, *app.Name)
	err = os.RemoveAll(metaPath)
	if err != nil {
		return err
	}

	return nil
}

func (service *appService) GetConfig(appName string) (*models.AppConfig, error) {
	metaYml, err := service.GetAppMeta(appName)
	if err != nil {
		return nil, err
	}
	return mappers.App.ToConfig(metaYml), nil
}

func (service *appService) UpdateConfig(appName string, newAppConfig *models.AppConfig) (*models.AppConfig, error) {
	existingMeta, err := service.GetAppMeta(appName)
	if err != nil {
		return nil, err
	}

	updatedMeta := existingMeta
	updatedMeta.TargetPorts = shared.ConvertInt64ArrayToUInt32Array(newAppConfig.TargetPorts)
	updatedMeta.Networks = newAppConfig.Networks

	err = service.WriteAppMeta(updatedMeta)
	if err != nil {
		return nil, err
	}

	return newAppConfig, nil
}
