package auth

import (
	"bimba/models"
	"bimba/utils"
	"errors"
)

var (
	ErrInvalidEmail    = errors.New("Email invalid")
	ErrInvalidPassword = errors.New("Password invalid")
)

func Signin(email, password string) (models.User, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, ErrInvalidEmail
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}
