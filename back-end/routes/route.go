package routes

import "github.com/gofiber/fiber/v2"

func Routes(router *fiber.App) {

	apiGroup := router.Group("/api")

	signUp(apiGroup)
}
