package controllers

import (
	"X-Sky/models"
	"X-Sky/services"

	"github.com/gofiber/fiber/v2"
)

func SignUp(router *fiber.Ctx) error {

	signData := models.SignUp{}

	if err := router.BodyParser(&signData); err != nil {
		return err
	}

	services.SignUp(signData)

	return nil
}
