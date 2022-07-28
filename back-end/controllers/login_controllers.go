package controllers

import (
	"X-Sky/models"
	"X-Sky/services"

	"github.com/gofiber/fiber/v2"
)

func Login(router *fiber.Ctx) error {

	loginData := models.Login{}

	if err := router.BodyParser(&loginData); err != nil {
		return err
	}

	err := services.Login(loginData)

	if err != nil {
		router.Redirect("http://localhost:5500/login")
		return err
	}

	router.Redirect("http://localhost:5500")
	return nil
}

func SignUp(router *fiber.Ctx) error {

	signData := models.SignUp{}

	if err := router.BodyParser(&signData); err != nil {
		return err
	}

	services.SignUp(signData)

	return nil
}
