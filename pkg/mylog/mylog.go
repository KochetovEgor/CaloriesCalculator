// mylog is a package for wrapping "log/slog" logger, for putting loggers in context.Context
// and from getting loggers from context.Context
package mylog

import (
	"context"
	"io"
	"log/slog"
)

// InitLogger initializes logger, that writes in w, and makes it default in "log/slog" package
func InitLogger(w io.Writer) {
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(handler))
}

type contextKey int

const loggerKey contextKey = 0

// NewContext putting logger in ctx and returns new context.Context
func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext gets logger from ctx
func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if ok {
		return logger
	}
	return slog.Default()
}
