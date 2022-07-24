package main

import (
	"X-Sky/config"
	middleware "X-Sky/middlewares"
	"X-Sky/routes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Define Fiber config & app.
	appCfg := config.LoadApp()
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)

	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	routes.Routes(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		fmt.Println("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		fmt.Println("Oops... server is not running! error:", err)
	}
}
