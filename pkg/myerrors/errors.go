// myerrors is a package, that makes my own Join(errs...) function (it is similar to errors.Join).
// And it also has some useful utils for error managment.
package myerrors

import (
	"strings"
)

const delimiter = "; "

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

// Join retunrs error, that wraps all given errs. Method Error of this new error returns string,
// that contains every wrapped error.Error with "; " delimiter.
func Join(errs ...error) error {
	var b strings.Builder
	for i, e := range errs {
		if e == nil {
			b.WriteString("")
		} else {
			b.WriteString(e.Error())
		}
		if i != len(errs)-1 {
			b.WriteString(delimiter)
		}
	}
	return &errorWrapper{msg: b.String(), wrappedErrors: errs}
}

// ExtractWrapped checks if err has method Unwrap() []error, and retunrs string slice
// with string values of all wrapped errors. It ignors nil errors. If err has not Unwrap() method,
// fucntion returns string slice, that contains only this err string value.
func ExtractWrapped(err error) []string {
	if err == nil {
		return nil
	}
	errSlice := []string{err.Error()}
	if wrappedErr, ok := err.(interface{ Unwrap() []error }); ok && wrappedErr != nil {
		errs := wrappedErr.Unwrap()
		errSlice = make([]string, 0, len(errs))
		for _, e := range errs {
			if e != nil {
				errSlice = append(errSlice, e.Error())
			}
		}
	}
	return errSlice
}
