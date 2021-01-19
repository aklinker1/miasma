package services

import (
	"fmt"
	"os"
)

type fileService struct{}

var Files = &fileService{}

// Utils

func (service *fileService) dirExists(path string) (bool, error) {
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

func (service *fileService) fileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("%s does not exist", path)
	}
	if stat.IsDir() {
		return false, fmt.Errorf("%s is a directory, not a file", path)
	}
	return true, nil
}
