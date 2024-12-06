package handler

import "testing"

func TestHandlerMain(t *testing.T) {
	lg.Debug("This is a debug message")
	lg.Fatal("This is a fatal  message")
	lg.Info("This is a info  message")
}
