package app

import (
	"openidea-shopyfyx/controller/user_controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	userController := user_controller.New()

	userGroup := app.Group("/v1/user")
	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)
}
