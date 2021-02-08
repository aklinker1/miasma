package editor

import (
	"errors"
	"os/exec"
)

func getEditorUtil() (path string, err error) {
	if path, err = exec.LookPath("vim"); err == nil {
		return path, nil
	}
	if path, err = exec.LookPath("nano"); err == nil {
		return path, nil
	}
	if path, err = exec.LookPath("vi"); err == nil {
		return path, nil
	}
	return "", errors.New("Could not find vim, nano, or vi in your $PATH")

}
