package auth

import (
	"unicode"
)

func isValidUsername(username string) error {
	if len(username) < 4 || len(username) > 50 {
		return &InvalidUsernameLengthError{}
	}
	for _, r := range username {
		if !(unicode.IsNumber(r) || unicode.IsLetter(r)) {
			return &InvalidUsernameContentError{}
		}
	}
	return nil
}

func isValidPassword(password string) error {
	if len(password) < 8 || len(password) > 50 {
		return &InvalidPasswordLengthError{}
	}

	wasUpper := false
	for _, r := range password {
		if !(unicode.IsNumber(r) || unicode.IsLetter(r)) {
			return &InvalidPasswordContentError{}
		}
		if unicode.IsUpper(r) {
			wasUpper = true
		}
	}

	if !wasUpper {
		return &InvalidPasswordContentError{}
	}

	return nil
}

func isValidCredentials(username, password string) error {
	if err := isValidUsername(username); err != nil {
		return err
	}
	if err := isValidPassword(username); err != nil {
		return err
	}
	return nil
}
