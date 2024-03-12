package app

import (
	"openidea-shopyfyx/controller/user_controller"
	"openidea-shopyfyx/db"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/service/user_service"
	"openidea-shopyfyx/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {

	dbPool, err := db.InitDBPool()
	utils.PanicErr(err)

	userRepository := user_repository.New(dbPool)
	userService := user_service.New(userRepository)
	userController := user_controller.New(userService)

	userGroup := app.Group("/v1/user")
	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)
}
