package myerrors

import (
	"strings"
)

type errorWrapper struct {
	msg           string
	wrappedErrors []error
}

func (e *errorWrapper) Error() string {
	return e.msg
}

func (e *errorWrapper) Unwrap() []error {
	return e.wrappedErrors
}

func Join(errs ...error) error {
	var b strings.Builder
	for i, e := range errs {
		b.WriteString(e.Error())
		if i != len(errs)-1 {
			b.WriteString("; ")
		}
	}
	return &errorWrapper{msg: b.String(), wrappedErrors: errs}
}

func ExtractWrapped(err error) []string {
	errSlice := []string{err.Error()}
	if wrappedErr, ok := err.(interface{ Unwrap() []error }); ok && wrappedErr != nil {
		errs := wrappedErr.Unwrap()
		errSlice = make([]string, len(errs))
		for i, e := range errs {
			errSlice[i] = e.Error()
		}
	}
	return errSlice
}
