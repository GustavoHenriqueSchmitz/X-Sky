package controllers

import (
	"X-Sky/models"
	"X-Sky/services"

	"github.com/gofiber/fiber/v2"
)

func Login(router *fiber.Ctx) error {

	loginData := models.Login{}

	if err := router.BodyParser(&loginData); err != nil {
		router.Redirect("http://localhost:3000/login")
		return nil
	}

	err := services.Login(loginData)

	if err != nil {
		router.Redirect("http://localhost:3000/login")
		return nil
	}

	router.Redirect("http://localhost:3000")
	return nil
}

func SignUp(router *fiber.Ctx) error {

	signData := models.SignUp{}

	if err := router.BodyParser(&signData); err != nil {
		router.Redirect("http://localhost:3000/sign-up")
		return nil
	}

	err := services.SignUp(signData)

	if err != nil {
		router.Redirect("http://localhost:3000/sign-up")
		return nil
	}

	router.Redirect("http://localhost:3000/login")
	return nil
}
