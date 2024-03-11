package app

import (
	"openidea-shopyfyx/controller/user_controller"
	"openidea-shopyfyx/service/user_service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	userService := user_service.New()
	userController := user_controller.New(userService)

	userGroup := app.Group("/v1/user")
	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)
}
