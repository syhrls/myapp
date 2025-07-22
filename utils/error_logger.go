package utils

import (
	"log"
	"path/filepath"
	"runtime"
	"time"
)

func LogError(err error) {
	if err == nil {
		return
	}
	pc, fullPath, line, ok := runtime.Caller(1)
	if !ok {
		log.Printf("%s [ERROR] (no caller info) -> %v\n", time.Now().Format(time.RFC3339), err)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	filename := filepath.Base(fullPath)

	log.Printf("[ERROR] %s:%d in %s() -> %v\n", filename, line, funcName, err)
}
