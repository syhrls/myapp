package utils

import (
	"fmt"
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
		fmt.Printf("%s [ERROR] (no caller info) -> %v\n", time.Now().Format(time.RFC3339), err)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	filename := filepath.Base(fullPath)
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("%s [ERROR] %s:%d in %s() -> %v\n", timestamp, filename, line, funcName, err)
}
