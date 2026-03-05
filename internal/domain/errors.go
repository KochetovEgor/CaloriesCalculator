package domain

type Err struct {
	msg string
}

func (e *Err) Error() string {
	return e.msg
}

var (
	ErrUserAlreadyExists = &Err{msg: "user already exists"}
)
