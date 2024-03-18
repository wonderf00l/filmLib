package auth

import (
	"unicode"
)

func isValidUsername(username string) error {
	if len(username) < 4 || len(username) > 50 {
		return &invalidUsernameLengthError{}
	}
	for _, r := range username {
		if !(unicode.IsNumber(r) || unicode.IsLetter(r)) {
			return &invalidUsernameContentError{}
		}
	}
	return nil
}

func isValidPassword(password string) error {
	if len(password) < 8 || len(password) > 50 {
		return &invalidPasswordLengthError{pass: password}
	}

	wasUpper := false
	for _, r := range password {
		if !(unicode.IsNumber(r) || unicode.IsLetter(r)) {
			return &invalidPasswordContentError{pass: password}
		}
		if unicode.IsUpper(r) {
			wasUpper = true
		}
	}

	if !wasUpper {
		return &invalidPasswordContentError{pass: password}
	}

	return nil
}

func isValidCredentials(username, password string) error {
	if err := isValidUsername(username); err != nil {
		return err
	}
	if err := isValidPassword(password); err != nil {
		return err
	}
	return nil
}
