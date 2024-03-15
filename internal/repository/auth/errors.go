package auth

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type ProfileAlreadyExistsError struct{}

func (e *ProfileAlreadyExistsError) Error() string {
	return "such profile already exists"
}

func (e *ProfileAlreadyExistsError) Type() errPkg.Type {
	return errPkg.ErrAlreadyExists
}

type ProfileNotFoundError struct{}

func (e *ProfileNotFoundError) Error() string {
	return "such profile doesn't exist"
}

func (e *ProfileNotFoundError) Type() errPkg.Type {
	return errPkg.ErrNotFound
}

type UserNotAuthenticatedError struct{}

func (e *UserNotAuthenticatedError) Error() string {
	return "user session not found"
}

func (e *UserNotAuthenticatedError) Type() errPkg.Type {
	return errPkg.ErrNoAuth
}
