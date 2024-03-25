package actor

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type invalidNameError struct {
}

func (e *invalidNameError) Error() string {
	return "invalid actor name was provided: should be - len[3:50], letters or ' symbol"
}

func (e *invalidNameError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type invalidGenderError struct {
}

func (e *invalidGenderError) Error() string {
	return "invalid actor gender was provided - should be [male, female]"
}

func (e *invalidGenderError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
