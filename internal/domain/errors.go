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
	ErrUserNotExists         = &Err{msg: "user not exists"}

	ErrProductAlreadyExists = &Err{msg: "product already exists"}
	ErrProductNotExists     = &Err{msg: "product not exists"}

	ErrPaswordTooLong   = &Err{msg: "password is too long"}
	ErrUsernameTooShort = &Err{msg: "username is too short"}

	ErrInvalidAccessToken = &Err{msg: "invalid access token"}

	ErrInternal = &Err{msg: "internal error"}
)
