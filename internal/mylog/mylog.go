package mylog

import (
	"context"
	"io"
	"log/slog"
)

func InitLogger(w io.Writer) {
	var handler slog.Handler
	handler = slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(handler))
}

type contextKey int

const loggerKey contextKey = 0

func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if ok {
		return logger
	}
	return slog.Default()
}
