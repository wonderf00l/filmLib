package actor

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type actorIDnotProvidedError struct{}

func (e *actorIDnotProvidedError) Error() string {
	return "provide actor ID in JSON body for updating"
}

func (e *actorIDnotProvidedError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
