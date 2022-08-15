package routes

import (
	"X-Sky/controllers"

	"github.com/gofiber/fiber/v2"
)

// Function login, to initializa the login routes.
func login(router fiber.Router) {
	router.Post("/login", controllers.Login)    // Login route.
	router.Post("/sign_up", controllers.SignUp) // Sing-up route.
}
