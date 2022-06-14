package fileutils

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	fileName = "pleasedelete.txt"
	dirsPath = "dir1/dir2"
	rootDir  = "dir1"

	nonExistentDir = "/@test"
)

func init() {
	deleteTempFiles()
}

func deleteTempFiles() {
	_ = os.RemoveAll(fileName)
	_ = os.RemoveAll(rootDir)
}

func TestWriteAndReadFile(t *testing.T) {
	defer deleteTempFiles()
	content := []byte("This is the content")
	errWrite := WriteToFile(fileName, content)
	if errWrite != nil {
		require.FailNow(t, "errWrite", errWrite)
	}

	fileContents, errRead := ReadFile(fileName)
	if errRead != nil {
		require.FailNow(t, "errRead", errRead)
	}

	require.Equal(t, content, fileContents)
}

func TestCreateDirs(t *testing.T) {
	defer deleteTempFiles()
	require.Equal(t, false, DirExist(dirsPath))

	err := CreateDirs(dirsPath, true)
	if err != nil {
		require.FailNow(t, "err", err)
	}

	require.Equal(t, true, DirExist(dirsPath))

	errAlreadyExistDir := CreateDirs(dirsPath, true)
	require.NotEqual(t, nil, errAlreadyExistDir)
}

func TestDeleteDirs(t *testing.T) {
	defer deleteTempFiles()
	errCreateDir := CreateDirs(dirsPath, true)
	if errCreateDir != nil {
		require.FailNow(t, "errCreateDir", errCreateDir)
	}

	require.Equal(t, true, DirExist(dirsPath))

	err := DeleteDirs(dirsPath)
	if err != nil {
		require.FailNow(t, "err", err)
	}

	require.Equal(t, false, DirExist(dirsPath))

	errNonExistentDir := DeleteDirs(nonExistentDir)
	require.NotEqual(t, nil, errNonExistentDir)
}

func TestDeleteFile(t *testing.T) {
	const name = "test.txt"
	defer DeleteFile(name)
	errCreate := WriteToFile(name, []byte("deleteme"))
	if errCreate != nil {
		require.FailNow(t, "errCreate", errCreate)
	}

	require.Equal(t, true, FileExist(name))

	err := DeleteFile(name)
	if err != nil {
		require.FailNow(t, "err", err)
	}

	require.Equal(t, false, FileExist(name))

	errNonExistentFile := DeleteFile(name)
	require.NotEqual(t, nil, errNonExistentFile)
}

func TestFileExists(t *testing.T) {
	defer deleteTempFiles()
	require.Equal(t, false, FileExist(fileName))

	content := []byte("This is the content")
	errWrite := WriteToFile(fileName, content)
	if errWrite != nil {
		require.FailNow(t, "errWrite", errWrite)
	}

	require.Equal(t, true, FileExist(fileName))
}
