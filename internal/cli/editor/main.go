package editor

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

func EditText(text string) (string, error) {
	// Fail when working on CI
	if os.Getenv("CI") == "true" {
		return "", errors.New("Cannot open editor when CI=true, because it is interactive")
	}

	// Create temp file
	tempDir := os.TempDir()
	file, err := ioutil.TempFile(tempDir, "*")
	if err != nil {
		return "", err
	}
	path := file.Name()
	defer os.Remove(path)

	// Load in the text we want to edit
	_, err = file.WriteString(text)
	if err != nil {
		return "", err
	}

	// Close the file so the editor can open it
	if err = file.Close(); err != nil {
		return "", err
	}

	// Prepare the editor
	var executable string
	envEditor := os.Getenv("EDITOR")
	if envEditor != "" {
		executable, err = exec.LookPath(envEditor)
	} else {
		executable, err = getEditorUtil()
	}
	if err != nil {
		return "", nil
	}

	// Run the editor
	cmd := exec.Command(executable, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	// Get the new contents
	newText, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(newText), nil
}
