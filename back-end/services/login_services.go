package services

import (
	"X-Sky/database"
	"X-Sky/models"
	"errors"
)

func Login(loginData models.Login) error {

	var password string
	err := database.DB.Table("users").Select("password").Where("email = " + "'" + loginData.Email + "'").Scan(&password)

	if err.Error != nil {
		return err.Error
	}

	if password == "" {
		return errors.New("Email não cadastrado.")
	} else if loginData.Password != password {
		return errors.New("Senha incorreta!")
	}

	return nil
}

func SignUp(signData models.SignUp) error {

	database.DB.Table("users").Create(signData)

	return nil
}
