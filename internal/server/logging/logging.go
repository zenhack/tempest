package logging

import (
	"os"

	"golang.org/x/exp/slog"
)

func NewLogger() *slog.Logger {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	return slog.New(opts.NewTextHandler(os.Stdout))
}
