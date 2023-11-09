package utils

import (
	"log"
	"runtime"
)

func CatchRuntimeErrors(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("[error] %s %d %v \n", file, line, err)
}
