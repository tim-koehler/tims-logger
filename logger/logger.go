package logger

import (
	"log"
)

// Logger interface
type Logger interface {
	Log(msg interface{})
	LogAndExit(msg interface{})
}

type textLogger struct {
}

// NewLogger creates new object of TextLogger
func NewLogger() Logger {
	return &textLogger{}
}

// Log Prints msg
func (l *textLogger) Log(msg interface{}) {
	log.Println(msg)
}

// LogAndExit Prints msg and exits with error code
func (l *textLogger) LogAndExit(msg interface{}) {
	log.Fatalln(msg)
}
