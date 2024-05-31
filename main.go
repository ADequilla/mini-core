package main

import (
	"fmt"
	"log"
	"mini-core/middleware"
	"mini-core/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	middleware.CreateConnection()

	app := fiber.New()
	app.Use(recover.New())
	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Access-Control-Allow-Origin",
	}))

	// Declare & initialize logger
	app.Use(logger.New())
	routers.SetupPrivateRoutes(app)

	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		// Bind to all available IP addresses and the specified port
		err := app.Listen(fmt.Sprintf("192.168.120.220:%s", middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
