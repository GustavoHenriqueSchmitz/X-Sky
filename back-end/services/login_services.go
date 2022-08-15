package services

import (
	"X-Sky/database"
	"X-Sky/models"
	"errors"
)

func Login(loginData models.Login) error {

	// It locates the password through the email.
	var password string
	database.DB.QueryRow("select password from users where email = '" + loginData.Email + "'").Scan(&password)

	// Validates if the email of the request body is registered.
	// If the email is registered, check if the body password is the same as the validated email.
	if password == "" {
		return errors.New("Email não cadastrado.")
	} else if loginData.Password != password {
		return errors.New("Senha incorreta!")
	}

	return nil
}

func SignUp(signData models.SignUp) error {

	// Try to insert in database the sign-up datas.
	_, err := database.DB.Query("insert into users(name, last_name, email, password, area_code, phone, perfil_photo) values('" + signData.Name + "', '" + signData.LastName + "', '" + signData.Email + "', '" + signData.Password + "', '" + signData.AreaCode + "', '" + signData.Phone + "', '" + signData.PerfilPhoto + "')	")

	if err != nil {
		return errors.New("Error registering data in the database.")
	}

	return nil
}
