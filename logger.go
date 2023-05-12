package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Log info", data)
}

func Error(data interface{}) {
	fmt.Println("Log error", data)
}

func Warning(message string, params interface{}) {
	fmt.Println("Log warning", message, "params: ", params)
}
