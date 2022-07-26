package services

import (
	"X-Sky/database"
	"X-Sky/models"
)

func SignUp(signData models.SignUp) error {

	database.DB.Table("users").Create(signData)

	return nil
}
