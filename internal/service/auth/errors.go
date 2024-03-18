package auth

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type InvalidPasswordError struct{}

func (e *InvalidPasswordError) Error() string {
	return "invalid password has been provided"
}

func (e *InvalidPasswordError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
