package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel string

type Logger struct {
	LogLevel   LogLevel
	OutputDest *os.File
	ToConsole  bool
}

func (l *Logger) log(level LogLevel, message string) {
	if l.isGoodLog(level) {
		time := time.Now().Format("2023-05-20 20:04:05")
		level := string(level)

		logMessage := fmt.Sprintf("[%s] [%s] %s", time, level, message)

		if l.ToConsole {
			fmt.Println(logMessage)
		} else {
			fmt.Fprintln(l.OutputDest, logMessage)
		}
	}
}

func (l *Logger) Debug(message string) {
	l.log("DEBUG", message)
}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Warning(message string) {
	l.log("WARNING", message)
}

func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

func (l *Logger) isGoodLog(level LogLevel) bool {
	switch level {
	case "DEBUG":
		return l.LogLevel == "DEBUG" || l.LogLevel == "INFO" || l.LogLevel == "WARNING" || l.LogLevel == "ERROR"
	case "INFO":
		return l.LogLevel == "INFO" || l.LogLevel == "WARNING" || l.LogLevel == "ERROR"
	case "WARNING":
		return l.LogLevel == "WARNING" || l.LogLevel == "ERROR"
	case "ERROR":
		return l.LogLevel == "ERROR"
	default:
		return false
	}
}
