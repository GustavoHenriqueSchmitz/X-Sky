package controllers

import (
	"X-Sky/models"
	"X-Sky/services"

	"github.com/gofiber/fiber/v2"
)

func Login(router *fiber.Ctx) error {

	loginData := models.Login{}

	if err := router.BodyParser(&loginData); err != nil {
		return router.SendStatus(500)
	}

	err := services.Login(loginData)

	if err != nil {
		router.Redirect("http://localhost:5500/login")
		return router.Status(403).SendString(err.Error())
	}

	router.Redirect("http://localhost:5500")
	return router.SendStatus(302)
}

func SignUp(router *fiber.Ctx) error {

	signData := models.SignUp{}

	if err := router.BodyParser(&signData); err != nil {
		return router.SendStatus(500)
	}

	err := services.SignUp(signData)

	if err != nil {
		return router.Status(400).SendString(err.Error())
	}

	return router.Status(200).SendString("Success")
}
