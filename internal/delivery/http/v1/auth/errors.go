package auth

import (
	"fmt"

	errPkg "github.com/wonderf00l/filmLib/internal/errors"
)

const (
	usernameMinLen, usernameMaxLen = 4, 50
	passMinLen, passMaxLen         = 8, 50
)

type invalidPasswordLengthError struct {
	got int
}

func (e *invalidPasswordLengthError) Error() string {
	return fmt.Sprintf("invalid password length: need [%d:%d], got - %d", passMinLen, passMaxLen, e.got)
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

type invalidUsernameLengthError struct {
	got int
}

func (e *invalidUsernameLengthError) Error() string {
	return fmt.Sprintf("invalid username length: need [%d:%d], got - %d", usernameMinLen, usernameMaxLen, e.got)
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
	return fmt.Sprintf("%s - %s, passwords don't match", e.pass1, e.pass2)
}

func (e *passwordsDontMatchError) Type() errPkg.Type {
	return errPkg.ErrInvalidInput
}
