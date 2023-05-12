package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Log info", data)
}

func Error(message string, data interface{}) {
	fmt.Println("Log error", data, "message", message)
}

func Warning(message string, params interface{}) {
	fmt.Println("Log warning", message, "params: ", params)
}
