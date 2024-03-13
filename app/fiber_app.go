package app

import (
	"log"

	"openidea-shopyfyx/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func InitFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  config.IdleTimeout,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
	})

	app.Use(logger.New())

	RegisterRoute(app)

	applicationHost := viper.GetString("APP_HOST")
	applicationPort := viper.GetString("APP_PORT")

	err := app.Listen(applicationHost + ":" + applicationPort)
	log.Fatal(err)
}
