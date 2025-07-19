package utils

import (
	"fmt"
	"runtime"
)

func LogError(err error) {
	if err == nil {
		return
	}
	// Dapatkan informasi caller (fungsi pemanggil)
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Error: %v (no caller info)\n", err)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()

	fmt.Printf("[ERROR] %s:%d in %s() -> %v\n", file, line, funcName, err)
}
