package util

import (
	"testing"
)

func TestLoggerInfo(t *testing.T) {
	logger := NewLogger()
	logger.Debug("This is a info message")
}

func TestLoggerFatal(t *testing.T) {
	logger := NewLogger()
	logger.Fatal("This is a fatal message")
}

func TestLoggerDebug(t *testing.T) {
	logger := NewLogger()
	logger.Debug("This is a debug message")
}
