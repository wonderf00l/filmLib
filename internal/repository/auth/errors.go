package auth

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type ProfileAlreadyExistsError struct{}

func (e *ProfileAlreadyExistsError) Error() string {
	return "such profile already exists"
}

func (e *ProfileAlreadyExistsError) Type() errPkg.Type {
	return errPkg.ErrAlreadyExists
}
