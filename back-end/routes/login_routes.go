package routes

import (
	"X-Sky/controllers"

	"github.com/gofiber/fiber/v2"
)

// Function login, to initializa the login routes.
func login(router fiber.Router) {

	// Agroup routes by login.
	router = router.Group("/login")

	router.Post("/", controllers.Login)         // Login route.
	router.Post("/sign-up", controllers.SignUp) // Sing-up route.
}
