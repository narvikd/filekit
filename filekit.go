package filekit

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// WriteToFile writes content to path. (It uses 600 permissions)
func WriteToFile(path string, content []byte) error {
	const permissions = 0600
	errWrite := os.WriteFile(filepath.Clean(path), content, permissions)
	if errWrite != nil {
		return fmt.Errorf("couldn't create file: %s. Error: %w", path, errWrite)
	}
	return nil
}

// ReadFile read the file from path.
func ReadFile(path string) ([]byte, error) {
	f, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %s. Error: %w", path, err)
	}
	return f, nil
}

// CreateDirs creates dirs. (It uses 750 permissions)
func CreateDirs(path string, checkIfDirExist bool) error {
	const permissions = 0750

	if checkIfDirExist && DirExist(filepath.Clean(path)) {
		return errors.New("directory already exist")
	}

	errMkdir := os.MkdirAll(filepath.Clean(path), permissions)
	if errMkdir != nil {
		return fmt.Errorf("couldn't create directory: %s. Error: %w", path, errMkdir)
	}
	return nil
}

// DeleteDirs deletes dirs only if they exist
func DeleteDirs(path string) error {
	if !DirExist(filepath.Clean(path)) {
		return errors.New("directory doesn't exist")
	}

	errDelDir := os.RemoveAll(filepath.Clean(path))
	if errDelDir != nil {
		return fmt.Errorf("couldn't delete directory: %s. Error: %w", path, errDelDir)
	}
	return nil
}

// DeleteFile only if the file exist
func DeleteFile(path string) error {
	if !FileExist(filepath.Clean(path)) {
		return errors.New("file doesn't exist")
	}

	errDelDir := os.Remove(filepath.Clean(path))
	if errDelDir != nil {
		return fmt.Errorf("couldn't delete file: %s. Error: %w", path, errDelDir)
	}
	return nil
}

// DirExist returns whether a directory exist
func DirExist(path string) bool {
	info, errStat := os.Stat(filepath.Clean(path))
	if errStat != nil {
		return false
	}
	if !info.IsDir() {
		return false
	}
	return true
}

// FileExist returns whether a file exist.
func FileExist(path string) bool {
	_, errStat := os.Stat(filepath.Clean(path))
	return errStat == nil
}
