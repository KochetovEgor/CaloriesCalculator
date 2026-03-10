package domain

import "errors"

type Err struct {
	msg string
}

func (e *Err) Error() string {
	return e.msg
}

func ExtractErr(err error) (error, bool) {
	return errors.AsType[*Err](err)
}

var (
	ErrUserAlreadyExists     = &Err{msg: "user already exists"}
	ErrInvalidUserOrPassword = &Err{msg: "invalid user or password"}

	ErrPaswordTooLong = &Err{msg: "password is too long"}
)
