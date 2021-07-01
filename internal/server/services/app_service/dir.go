package app_service

import (
	"os"

	"github.com/aklinker1/miasma/internal/server/services/file_service"
	"github.com/aklinker1/miasma/internal/shared/log"
)

func AppsDir() (dir string, err error) {
	log.V("app_service.AppsDir()")
	dir = "/data/miasma/apps"
	if exists, _ := file_service.DirExists(dir); !exists {
		// https://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value
		err = os.MkdirAll(dir, 0755)
	}
	return dir, err
}
