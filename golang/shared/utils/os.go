package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

func GetCurrrentFolder() string {
	_, filename, _, _ := runtime.Caller(1)
	directory := filepath.Dir(filename)
	return directory
}

func ReadProjectFile(path string) string {
	dat, err := os.ReadFile(filepath.Join(GetCurrrentFolder(), "../../../", path))
	Check(err)
	return string(dat)
}
