package auth

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type profileAlreadyExistsError struct{}

func (e *profileAlreadyExistsError) Error() string {
	return "such profile already exists"
}

func (e *profileAlreadyExistsError) Type() errPkg.Type {
	return errPkg.ErrAlreadyExists
}

type profileNotFoundError struct{}

func (e *profileNotFoundError) Error() string {
	return "such profile doesn't exist"
}

func (e *profileNotFoundError) Type() errPkg.Type {
	return errPkg.ErrNotFound
}

type sessionNotFoundError struct{}

func (e *sessionNotFoundError) Error() string {
	return "user session wasn't found"
}

func (e *sessionNotFoundError) Type() errPkg.Type {
	return errPkg.ErrNoAuth
}
