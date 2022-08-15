package routes

import "github.com/gofiber/fiber/v2"

// Routes function, to group the routes in the api group, and initialize them.
func Routes(router *fiber.App) {

	// Agroup routes by api.
	apiGroup := router.Group("/api")

	login(apiGroup) // Login Routes
}
