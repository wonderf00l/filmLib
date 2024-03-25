package auth

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type invalidPasswordError struct{}

func (e *invalidPasswordError) Error() string {
	return "invalid password has been provided"
}

func (e *invalidPasswordError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
