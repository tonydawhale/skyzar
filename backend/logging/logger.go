package logging

import (
	"log"
	"time"

	"github.com/TwiN/go-color"
)

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Debug(message string) {
	println(color.With(color.Cyan, "[" + getTime() + "] - [Skyzar-DEBUG] " + message))
}

func LogWarn(message string) {
	println(color.With(color.Yellow, "[" + getTime() + "] - [Skyzar-WARN] " + message))
}
func Log(message string) {
	println(color.With(color.Green, "[" + getTime() + "] - [Skyzar] " + message))
}
func Error(message string) {
	println(color.With(color.Red, "[" + getTime() + "] - [Skyzar-ERROR] " + message))
}
func LogFatal(message string) {
	log.Fatal(color.With(color.Red, "[" + getTime() + "] - [Skyzar-FATAL] " + message))
}