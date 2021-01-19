package services

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aklinker1/miasma/internal/gen/models"
	"github.com/aklinker1/miasma/internal/utils/mappers"
	"github.com/aklinker1/miasma/internal/utils/types"

	"gopkg.in/yaml.v2"
)

type fileService struct{}

var Files = &fileService{}

// Utils

func (files *fileService) dirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Println(err)
		return false, fmt.Errorf("%s does not exist", path)
	}
	if !stat.IsDir() {
		return false, fmt.Errorf("%s is a file, not a directory", path)
	}
	return true, nil
}

func (files *fileService) fileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("%s does not exist", path)
	}
	if stat.IsDir() {
		return false, fmt.Errorf("%s is a directory, not a file", path)
	}
	return true, nil
}

// App

func (files *fileService) AppsDir() (dir string, err error) {
	dir = "/data/miasma/apps"
	if exists, _ := files.dirExists(dir); !exists {
		// https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value
		err = os.MkdirAll(dir, 0777)
	}
	return dir, err
}

func (files *fileService) ReadApp(appName string) (*models.App, error) {
	appsDir, err := files.AppsDir()
	if err != nil {
		return nil, err
	}
	appDir := fmt.Sprintf("%s/%s", appsDir, appName)
	if exists, _ := files.dirExists(appDir); !exists {
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

	return mappers.App.FromMeta(appName, metaYml), err
}

func (files *fileService) ReadApps(showHidden bool) ([]*models.App, error) {
	appsDir, err := files.AppsDir()
	if err != nil {
		return nil, err
	}
	appDirs, err := ioutil.ReadDir(appsDir)
	if err != nil {
		return nil, err
	}

	result := []*models.App{}
	for _, appDir := range appDirs {
		app, err := files.ReadApp(appDir.Name())
		if err != nil {
			return nil, err
		}
		if showHidden || !app.Hidden {
			result = append(result, app)
		}
	}
	return result, nil
}

func (files *fileService) GenerateAppComposeFile(appName string) {

}

func (files *fileService) ListApps() {

}
