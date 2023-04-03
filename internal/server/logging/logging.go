// Package logging adds helpers for logging
package logging

import (
	"os"

	"golang.org/x/exp/slog"
)

// Create a new logger with our preferred settings.
func NewLogger() *slog.Logger {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	return slog.New(opts.NewTextHandler(os.Stdout))
}

// Log a message and then panic
func Panic(l *slog.Logger, msg string, args ...any) {
	msg = "FATAL: " + msg
	l.Error(msg, args...)
	panic(msg)
}
