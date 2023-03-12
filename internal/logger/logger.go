package logger

import (
	"errors"
	"fmt"
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

func Fatal(msg string) error {
	return errors.New(msg)
}

func Fatalf(format string, args ...any) error {
	return fmt.Errorf(format+"\n", args...)
}
