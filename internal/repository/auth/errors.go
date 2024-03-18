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

type SessionNotFoundError struct{}

func (e *SessionNotFoundError) Error() string {
	return "no such session"
}

func (e *SessionNotFoundError) Type() errPkg.Type {
	return errPkg.ErrNoAuth
}
