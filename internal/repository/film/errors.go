package film

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type filmAlreadyExistsError struct{}

func (e *filmAlreadyExistsError) Error() string {
	return "such profile already exists"
}

func (e *filmAlreadyExistsError) Type() errPkg.Type {
	return errPkg.ErrAlreadyExists
}
