package user

import (
	"dashboard-api/utils"
	"errors"
)

//ValidateLogin :
func ValidateLogin(user Login) error {
	if len(user.Username) <= 0 {
		return errors.New(utils.UsernameError)
	}
	if len(user.Password) <= 0 {
		return errors.New(utils.PasswordError)
	}
	return nil
}