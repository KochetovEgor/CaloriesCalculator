package mylog

import (
	"context"
	"errors"
)

type Error struct {
	err   error
	attrs []any
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func WrapError(err error, attrs ...any) error {
	if attrErr, ok := errors.AsType[*Error](err); ok {
		attrs = append(attrs, attrErr.attrs...)
	}

	attrErr := &Error{err: err, attrs: attrs}
	return attrErr
}

const errorKey contextKey = "error"

func ErrToContext(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey, err)
}

func ErrFromContext(ctx context.Context) error {
	err, _ := ctx.Value(errorKey).(error)
	return err
}
