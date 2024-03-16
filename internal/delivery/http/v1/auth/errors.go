package auth

import (
	"fmt"

	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

type invalidPasswordLengthError struct {
	pass string
}

func (e *invalidPasswordLengthError) Error() string {
	return fmt.Sprintf("%s - invalid password length: need [8:50]\n", e.pass)
}

func (e *invalidPasswordLengthError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type invalidPasswordContentError struct {
	pass string
}

func (e *invalidPasswordContentError) Error() string {
	return fmt.Sprintf("%s - invalid password content: should be letters or numbers with at least one symbol in upper case", e.pass)
}

func (e *invalidPasswordContentError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type invalidUsernameLengthError struct{}

func (e *invalidUsernameLengthError) Error() string {
	return "invalid username length: need [4:50]"
}

func (e *invalidUsernameLengthError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type invalidUsernameContentError struct{}

func (e *invalidUsernameContentError) Error() string {
	return "invalid username content: should be letters or numbers"
}

func (e *invalidUsernameContentError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}

type noOldPasswordError struct{}

func (e *noOldPasswordError) Error() string {
	return "provide old password to change profile data"
}

func (e *noOldPasswordError) Type() errPkg.Type {
	return errPkg.ErrNoAccess
}

type passwordsDontMatchError struct {
	pass1, pass2 string
}

func (e *passwordsDontMatchError) Error() string {
	return fmt.Sprintf("%s - %s, passwords don't match\n", e.pass1, e.pass2)
}

func (e *passwordsDontMatchError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
