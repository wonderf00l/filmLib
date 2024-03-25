package actor

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type invalidTimeFormatError struct {
}

func (e *invalidTimeFormatError) Error() string {
	return "invalid time format, should be yyyy-mm-dd"
}

func (e *invalidTimeFormatError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
