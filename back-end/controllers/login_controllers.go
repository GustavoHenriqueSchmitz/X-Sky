package controllers

import (
	"X-Sky/models"
	"X-Sky/services"

	"github.com/gofiber/fiber/v2"
)

// Function Login, controller of login route.
func Login(router *fiber.Ctx) error {

	// Defines the data model, and takes the data from the request body.
	loginData := models.Login{}

	if err := router.BodyParser(&loginData); err != nil {
		router.Redirect("http://localhost:3000/login")
		return nil
	}

	// Calls the login service and returns the responses.
	err := services.Login(loginData)

	if err != nil {
		router.Redirect("http://localhost:3000/login")
		return nil
	}

	router.Redirect("http://localhost:3000")
	return nil
}

// Function SignUp, controller of sign-up route.
func SignUp(router *fiber.Ctx) error {

	// Defines the data model, and takes the data from the request body.
	signData := models.SignUp{}

	if err := router.BodyParser(&signData); err != nil {
		router.Redirect("http://localhost:3000/sign-up")
		return nil
	}

	// Calls the sign-up service and returns the responses.
	err := services.SignUp(signData)

	if err != nil {
		router.Redirect("http://localhost:3000/sign-up")
		return nil
	}

	router.Redirect("http://localhost:3000/login")
	return nil
}
