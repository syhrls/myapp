package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

func Info(message string) {
	fmt.Println(colorGreen + "[INFO] " + message + colorReset)
}

func Debug(message string) {
	fmt.Println(colorCyan + "[DEBUG] " + message + colorReset)
}

func Error(message string) {
	if message == "" {
		return
	}

	pc, fullPath, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf(colorYellow+"[ERROR] (no caller info) -> %v"+colorReset+"\n", message)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()

	wd, _ := os.Getwd()
	relPath, err := filepath.Rel(wd, fullPath)
	if err != nil {
		relPath = fullPath
	}

	fmt.Printf(colorYellow+"[ERROR] %s:%d in %s() -> %v"+colorReset+"\n", relPath, line, funcName, message)
}

func Fatal(message string) {
	fmt.Print(colorRed + "[FATAL] " + message + colorReset)
}

func InfoWithContext(c *gin.Context, message string) {
	reqID, _ := c.Get("request_id")
	fmt.Printf(colorGreen+"[INFO] [%v] %v"+colorReset+"\n", reqID, message)
}

type ColorWriter struct {
	Writer io.Writer
}

func (cw ColorWriter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("[GIN-debug]")) {
		colored := colorGreen + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	}else if bytes.Contains(p, []byte("[ERROR]")) {
		colored := colorRed + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	} else if bytes.Contains(p, []byte("[GIN]")) {
		colored := colorGreen + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	} else if bytes.Contains(p, []byte("[WARN]")) {
		colored := colorYellow + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	} else if bytes.Contains(p, []byte("[DEBUG]")) {
		colored := colorCyan + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	}
	return cw.Writer.Write(p)
}