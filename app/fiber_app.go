package app

import (
	"log"

	"openidea-shopyfyx/config"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func InitFiberApp() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  config.IdleTimeout,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
	})

	RegisterRoute(app)

	viper.BindEnv("APP_HOST")
	viper.BindEnv("APP_PORT")

	applicationHost := viper.GetString("APP_HOST")
	applicationPort := viper.GetString("APP_PORT")

	err := app.Listen(applicationHost + ":" + applicationPort)
	log.Fatal(err)
}
