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
		return router.Status(500).JSON(fiber.Map{
			"message": err,
			"error":   true,
		})
	}

	// Calls the login service and returns the responses.
	err := services.Login(loginData)

	if err != nil {
		return router.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return router.Status(200).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
	})
}

// Function SignUp, controller of sign-up route.
func SignUp(router *fiber.Ctx) error {

	// Defines the data model, and takes the data from the request body.
	signData := models.SignUp{}

	if err := router.BodyParser(&signData); err != nil {
		return router.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	// Calls the sign-up service and returns the responses.
	err := services.SignUp(signData)

	if err != nil {
		return router.Status(400).JSON(fiber.Map{
			"message": "Error when trying to register data in the database.",
			"error":   true,
		})
	}

	return router.Status(201).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
	})
}
