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
    loc *time.Location
}

func (w JakartaLogWriter) Write(p []byte) (n int, err error) {
    timestamp := time.Now().In(w.loc).Format("2006/01/02 15:04:05")
    newLog := append([]byte(timestamp+" "), p...)
    return w.out.Write(newLog)
}

func InitLoggerWIB() {
    loc, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        log.Println("[WARN] Falling back to manual +07 timezone due to missing Asia/Jakarta")
        loc = time.FixedZone("WIB", 7*3600)
    }

    log.SetFlags(0)
    log.SetOutput(JakartaLogWriter{
        out: os.Stdout,
        loc: loc,
    })
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