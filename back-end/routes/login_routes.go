package routes

import (
	"X-Sky/controllers"

	"github.com/gofiber/fiber/v2"
)

func signUp(router fiber.Router) {

	router = router.Group("/login")

	router.Put("/sign-up", controllers.SignUp)
}
