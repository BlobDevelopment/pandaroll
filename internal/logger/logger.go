package logger

import (
	"fmt"
	"os"
)

func Info(msg string) {
	fmt.Println("[INFO] " + msg)
}

func Infof(format string, args ...any) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

func Error(msg string) {
	fmt.Println("[ERROR] " + msg)
}

func Errorf(format string, args ...any) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

func Fatal(msg string) {
	Error(msg)
	os.Exit(1)
}

func Fatalf(format string, args ...any) {
	Errorf(format, args...)
	os.Exit(1)
}
