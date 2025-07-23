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
		Warn("Falling back to manual +07 timezone due to missing Asia/Jakarta")
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

	colorBlack  = "\033[30m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func Success(message string) {
	log.Println(colorGreen + "[SUCCESS] " + message + colorReset)
}

func Info(message string) {
	log.Println(colorBlue + "[INFO] " + message + colorReset)
}

func Debug(message string) {
	log.Println(colorPurple + "[DEBUG] " + message + colorReset)
}

func Warn(message string) {
	log.Println(colorYellow + "[WARN] " + message + colorReset)
}

func Error(message string) {
	if message == "" {
		return
	}

	pc, fullPath, line, ok := runtime.Caller(2)
	if !ok {
		log.Printf(colorRed+"[ERROR] (no caller info) -> %v"+colorReset+"\n", message)
		return
	}

	funcName := runtime.FuncForPC(pc).Name()

	wd, _ := os.Getwd()
	relPath, err := filepath.Rel(wd, fullPath)
	if err != nil {
		relPath = fullPath
	}

	log.Printf(colorRed+"[ERROR] %s:%d in %s() -> %v"+colorReset+"\n", relPath, line, funcName, message)
}

func Fatal(message string) {
	log.Print(colorRed + "[FATAL] " + message + colorReset)
}

func InfoWithContext(c *gin.Context, format string, args ...any) {
    requestID, exists := c.Get("request_id")
    if !exists {
        requestID = "unknown"
    }

    // Hanya cetak 1 timestamp manual
    Info(fmt.Sprintf("[REQUEST ID: %s] %s", requestID, fmt.Sprintf(format, args...)))
}

type ColorWriter struct {
	Writer io.Writer
}

func (cw ColorWriter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("[GIN-debug]")) {
		colored := colorBlue + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	}else if bytes.Contains(p, []byte("[ERROR]")) {
		colored := colorRed + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	} else if bytes.Contains(p, []byte("[GIN]")) {
		colored := colorBlue + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	} else if bytes.Contains(p, []byte("[WARN]")) {
		colored := colorYellow + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	} else if bytes.Contains(p, []byte("[DEBUG]")) {
		colored := colorPurple + string(p) + colorReset
		return cw.Writer.Write([]byte(colored))
	}
	return cw.Writer.Write(p)
}

func LogStartEnd() gin.HandlerFunc {
	return func(c *gin.Context) {
		InfoWithContext(c, "Start processing %s", c.Request.URL.Path)
		c.Next()
		InfoWithContext(c, "End processing %s", c.Request.URL.Path)
	}
}