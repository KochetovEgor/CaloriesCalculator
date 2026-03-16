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
	// User errors
	ErrUserAlreadyExists     = &Err{msg: "user already exists"}
	ErrInvalidUserOrPassword = &Err{msg: "invalid user or password"}
	ErrUserNotExists         = &Err{msg: "user not exists"}

	// Validation User errors
	ErrPaswordTooLong   = &Err{msg: "password is too long"}
	ErrUsernameTooShort = &Err{msg: "username is too short"}

	// Product errors
	ErrProductAlreadyExists = &Err{msg: "product already exists"}
	ErrProductNotExists     = &Err{msg: "product not exists"}

	//Validation Product errors
	ErrBaseWeightMustBePositive    = &Err{msg: "base weight must be positive"}
	ErrBasePortionMustBePositive   = &Err{msg: "base portion must be positive"}
	ErrFatsMustBePositive          = &Err{msg: "fats must be positive"}
	ErrProteinsMustBePositive      = &Err{msg: "proteins must be positive"}
	ErrCarbohydratesMustBePositive = &Err{msg: "carbohydrates must be positive"}

	// Other errors
	ErrInvalidAccessToken = &Err{msg: "invalid access token"}
	ErrInternal           = &Err{msg: "internal error"}
)
