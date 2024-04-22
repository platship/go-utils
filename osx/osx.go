package osx

import (
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
)

func CreateDirIsNotExist(dir string, perm os.FileMode) (err error) {
	if fileutil.IsExist(dir) {
		return
	}
	return os.MkdirAll(dir, perm)
}

func CreateFileWithDir(path string, content []byte, perm os.FileMode) (err error) {
	if err = CreateDirIsNotExist(filepath.Dir(path), perm); err != nil {
		return
	}
	return CreateFile(path, content)
}

func CreateFile(path string, content []byte) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	if len(content) > 0 {
		_, err = file.Write(content)
	}
	return
}
