package logger

import (
	"log"
)

type Logger interface {
	Log(msg interface{})
}

type TextLogger struct {
}

func NewLogger() Logger {
	return &TextLogger{}
}

func (l *TextLogger) Log(msg interface{}) {
	log.Println(msg)
}
