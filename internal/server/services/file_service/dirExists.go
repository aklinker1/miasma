package file_service

import (
	"fmt"
	"os"

	"github.com/aklinker1/miasma/internal/shared/log"
)

func DirExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.V("%v", err)
		return false, fmt.Errorf("%s does not exist", path)
	}
	if !stat.IsDir() {
		return false, fmt.Errorf("%s is a file, not a directory", path)
	}
	return true, nil
}
