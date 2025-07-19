package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func LogError(err error) {
	if err == nil {
		return
	}
	pc, fullPath, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error: %v (no caller info)\n", err)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	filename := filepath.Base(fullPath)

	fmt.Printf("[ERROR] %s:%d in %s() -> %v\n", filename, line, funcName, err)
}
