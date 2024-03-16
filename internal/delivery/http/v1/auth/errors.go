package auth

import errPkg "github.com/wonderf00l/filmLib/internal/errors"

type InvalidPasswordLengthError struct{}

func (e *InvalidPasswordLengthError) Error() string {
	return "invalid password length: need [8:50]"
}

func (e *InvalidPasswordLengthError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type InvalidPasswordContentError struct{}

func (e *InvalidPasswordContentError) Error() string {
	return "invalid password content: should be letters or numbers with at least one symbol in upper case"
}

func (e *InvalidPasswordContentError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type InvalidUsernameLengthError struct{}

func (e *InvalidUsernameLengthError) Error() string {
	return "invalid username length: need [4:50]"
}

func (e *InvalidUsernameLengthError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type InvalidUsernameContentError struct{}

func (e *InvalidUsernameContentError) Error() string {
	return "invalid username content: should be letters or numbers"
}

func (e *InvalidUsernameContentError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
