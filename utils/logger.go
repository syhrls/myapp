package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type JakartaLogWriter struct {
    out io.Writer
}

func (w JakartaLogWriter) Write(p []byte) (n int, err error) {
    loc, _ := time.LoadLocation("Asia/Jakarta")
    timestamp := time.Now().In(loc).Format("2006/01/02 15:04:05")

    // Tambahkan waktu Jakarta + pesan asli
    newLog := append([]byte(timestamp+" "), p...)
    return w.out.Write(newLog)
}

// Inisialisasi logger agar pakai zona waktu Asia/Jakarta
func InitLoggerWIB() {
    log.SetFlags(0) // hilangkan waktu default
    log.SetOutput(JakartaLogWriter{out: os.Stdout})
}

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
	log.Print(colorRed + "[FATAL] " + message + colorReset)
}

func InfoWithContext(c *gin.Context, format string, args ...any) {
    requestID, exists := c.Get("request_id")
    if !exists {
        requestID = "unknown"
    }


    // Waktu zona +7 / Asia/Jakarta

    // Hanya cetak 1 timestamp manual
    log.Printf(" [INFO] [%s] %s", requestID, fmt.Sprintf(format, args...))
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