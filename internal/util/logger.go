package util

import (
	"log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Fatal(v ...interface{})
}

type ConsoleLogger struct {
	debug *log.Logger
	info  *log.Logger
	fatal *log.Logger
}

func NewLogger() *ConsoleLogger {
	return &ConsoleLogger{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		fatal: log.New(os.Stdout, "FATAL: ", log.LstdFlags),
	}
}

func (l *ConsoleLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *ConsoleLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *ConsoleLogger) Fatal(v ...interface{}) {
	l.fatal.Println(v...)
}
