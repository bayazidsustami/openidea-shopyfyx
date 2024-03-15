package app

import (
	"log"

	"openidea-shopyfyx/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  config.IdleTimeout,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
	})

	app.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	))
	app.Use(logger.New())

	RegisterRoute(app)

	err := app.Listen("localhost:8000")
	log.Fatal(err)
}
