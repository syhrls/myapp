package utils

import (
	"log"
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
	log.Println(colorGreen + "[INFO] " + message + colorReset)
}

func Debug(message string) {
	log.Println(colorCyan + "[DEBUG] " + message + colorReset)
}

func Error(message string) {
    if message == "" {
        return
    }

    pc, fullPath, line, ok := runtime.Caller(2)
    if !ok {
        log.Printf(colorYellow+"[ERROR] (no caller info) -> %v"+colorReset+"\n", message)
        return
    }

    funcName := runtime.FuncForPC(pc).Name()

    wd, _ := os.Getwd()
    relPath, err := filepath.Rel(wd, fullPath)
    if err != nil {
        relPath = fullPath
    }

    log.Printf(colorYellow+"[ERROR] %s:%d in %s() -> %v"+colorReset+"\n", relPath, line, funcName, message)
}

func Fatal(message string) {
	log.Fatalln(colorRed + "[FATAL] " + message + colorReset)
}

func InfoWithContext(c *gin.Context, message string) {
	reqID, _ := c.Get("request_id")
	log.Printf(colorYellow+"[INFO] [%v] %v"+colorReset+"\n", reqID, message)
}
