package utils

import (
	"log"
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
	log.Println(colorYellow + "[ERROR] " + message + colorReset)
}

func Fatal(message string) {
	log.Fatalln(colorRed + "[FATAL] " + message + colorReset)
}
