// mylog is a package for wrapping "log/slog" logger, for putting loggers in context.Context
// and from getting loggers from context.Context
package mylog

import (
	"context"
	"errors"
	"io"
	"log/slog"
)

type errorMiddleware struct {
	next slog.Handler
}

func newErrorMiddleware(next slog.Handler) *errorMiddleware {
	return &errorMiddleware{next: next}
}

func (em *errorMiddleware) Enabled(ctx context.Context, level slog.Level) bool {
	return em.next.Enabled(ctx, level)
}

func (em *errorMiddleware) Handle(ctx context.Context, record slog.Record) error {
	if record.Level == slog.LevelError {
		err := ErrFromContext(ctx)
		if attrsErr, ok := errors.AsType[*Error](err); ok {
			record.Add(attrsErr.attrs...)
		}
	}

	return em.next.Handle(ctx, record)
}

func (em *errorMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &errorMiddleware{next: em.next.WithAttrs(attrs)}
}

func (em *errorMiddleware) WithGroup(name string) slog.Handler {
	return &errorMiddleware{next: em.next.WithGroup(name)}
}

// InitLogger initializes logger, that writes in w, and makes it default in "log/slog" package
func InitLogger(w io.Writer, level slog.Level) {
	handlerJSON := slog.NewJSONHandler(w, &slog.HandlerOptions{Level: level})
	handler := newErrorMiddleware(handlerJSON)
	slog.SetDefault(slog.New(handler))
}

type contextKey string

const loggerKey contextKey = "log"

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
