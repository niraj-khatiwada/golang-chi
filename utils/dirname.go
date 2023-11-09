package utils

import (
	"path"
	"runtime"
)

// GetDir = Returns the name of the current directory
func GetDir() string {
	_, file, _, _ := runtime.Caller(1)
	return path.Dir(file)
}
