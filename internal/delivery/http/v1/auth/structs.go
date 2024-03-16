package auth

import entity "github.com/wonderf00l/filmLib/internal/entity/auth"

type signupData struct {
	Username  *string `json:"username" example:"newbie123"`
	Password  *string `json:"password" example:"helloWorld"`
	RoleToken *string `json:"role_token" example:"123"`
}

func (d *signupData) Validate() error {
	if d.Username == nil {
		return &invalidUsernameLengthError{}
	}
	if d.Password == nil {
		return &invalidPasswordLengthError{pass: ""}
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
		return &invalidUsernameLengthError{}
	}
	if d.Password == nil {
		return &invalidPasswordLengthError{}
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

type updateData struct {
	NewUsername         *string `json:"new_username"`
	OldPassword         *string `json:"old_password"`
	NewPassword         *string `json:"new_password"`
	RepeatedNewPassword *string `json:"new_password_repeated"`
	NewRoleToken        *string `json:"new_role_token"`
}

func (d *updateData) Validate() error {
	if d.OldPassword == nil {
		return &noOldPasswordError{}
	}

	if err := isValidPassword(*d.OldPassword); err != nil {
		return err
	}

	if d.NewUsername != nil {
		if err := isValidUsername(*d.NewUsername); err != nil {
			return err
		}
	}

	if d.NewPassword != nil {
		if err := isValidPassword(*d.NewPassword); err != nil {
			return err
		}
	}

	if d.NewPassword != nil && d.RepeatedNewPassword == nil {
		return &invalidPasswordLengthError{pass: ""}
	}

	if d.NewPassword != nil && *d.NewPassword != *d.RepeatedNewPassword {
		return &passwordsDontMatchError{*d.NewPassword, *d.RepeatedNewPassword}
	}

	return nil
}

func updateDataDeliveryToService(data updateData) (entity.Profile, string) {
	tok := ""
	if data.NewRoleToken != nil {
		tok = *data.NewRoleToken
	}

	newUsername := ""
	if data.NewUsername != nil {
		newUsername = *data.NewUsername
	}

	newPass := ""
	if data.NewPassword != nil {
		newPass = *data.NewPassword
	}

	return entity.Profile{
		Username: newUsername,
		Password: newPass,
	}, tok
}
