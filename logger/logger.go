package logger

import (
	"fmt"
)

func Info(event string, details string, rest ...interface{}) {

	fmt.Println("INFO - event:", event, "details:", details, rest)
}

func Error(event string, details string, err string, rest ...interface{}) {
	fmt.Println("INFO - event:", event, "details:", details, "error:", err, rest)
}
