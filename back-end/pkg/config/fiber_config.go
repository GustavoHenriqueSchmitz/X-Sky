package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
func FiberConfig() fiber.Config {

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(App{}.ReadTimeout),
	}
}
