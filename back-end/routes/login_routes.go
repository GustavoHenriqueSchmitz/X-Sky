package routes

import (
	"X-Sky/controllers"

	"github.com/gofiber/fiber/v2"
)

func signUp(router fiber.Router) {

	router.Post("/login", controllers.Login)
	router.Post("/sign_up", controllers.SignUp)
}
