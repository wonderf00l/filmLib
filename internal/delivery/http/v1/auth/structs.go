package auth

import entity "github.com/wonderf00l/filmLib/internal/entity/auth"

type signupData struct {
	Username  *string `json:"username" example:"newbie123"`
	Password  *string `json:"password" example:"helloWorld"`
	RoleToken *string `json:"role_token" example:"admToken"`
}

func (d *signupData) Validate() error {
	if d.Username == nil {
		return &InvalidUsernameLengthError{}
	}
	if d.Password == nil {
		return &InvalidPasswordLengthError{}
	}

	if err := isValidCredentials(*d.Username, *d.Password); err != nil {
		return err
	}

	return nil
}

func signupDataDeliveryToService(data signupData) (entity.Profile, string) {
	tok := ""
	if data.RoleToken != nil {
		tok = *data.RoleToken
	}
	return entity.Profile{
		Username: *data.Username,
		Password: *data.Password,
	}, tok
}

type loginData struct {
	Username *string `json:"username" example:"newbie123"`
	Password *string `json:"password" example:"helloWorld"`
}

func (d *loginData) Validate() error {
	if d.Username == nil {
		return &InvalidUsernameLengthError{}
	}
	if d.Password == nil {
		return &InvalidPasswordLengthError{}
	}

	if err := isValidCredentials(*d.Username, *d.Password); err != nil {
		return err
	}

	return nil
}

func loginDataDeliveryToService(data loginData) entity.Profile {
	return entity.Profile{
		Username: *data.Username,
		Password: *data.Password,
	}
}
