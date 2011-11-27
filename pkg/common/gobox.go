package common

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func getRelativeGoboxBinaryPath() (string, error) {
	callname := os.Args[0]
	// First check: Is gobox in $PATH?
	path, e := exec.LookPath(callname)
	if e == nil {
		return path, nil
	}

	// Second check: Is gobox in the current directory?
	cwd, e := os.Getwd()
	if e != nil {
		return "", e
	}
	path = filepath.Join(cwd, "gobox")
	if PathExists(path) {
		return path, nil
	}
	return "", errors.New("Could not find gobox binary")
}

func GetGoboxBinaryPath() (string, error) {
	relpath, e := getRelativeGoboxBinaryPath()
	if e != nil {
		return "", e
	}
	return filepath.Abs(relpath)
}
