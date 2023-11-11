package utils

import (
	"path"
	"runtime"
)

// GetCurrentDir = Returns the name of the current directory
func GetCurrentDir() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Dir(file)
}
