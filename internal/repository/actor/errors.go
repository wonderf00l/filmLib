package actor

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type actorAlreadyExistsError struct{}

func (e *actorAlreadyExistsError) Error() string {
	return "such actor already exists"
}

func (e *actorAlreadyExistsError) Type() errPkg.Type {
	return errPkg.ErrAlreadyExists
}

type actorNotFoundError struct{}

func (e *actorNotFoundError) Error() string {
	return "such actor doesn't exist"
}

func (e *actorNotFoundError) Type() errPkg.Type {
	return errPkg.ErrNotFound
}
