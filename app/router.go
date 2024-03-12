package app

import (
	"openidea-shopyfyx/controller/user_controller"
	"openidea-shopyfyx/db"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/user_service"
	"openidea-shopyfyx/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {

	validator := validator.New()
	dbPool, err := db.InitDBPool()
	utils.PanicErr(err)

	authService := auth_service.New()

	userRepository := user_repository.New(dbPool)
	userService := user_service.New(userRepository, validator, authService)
	userController := user_controller.New(userService)

	userGroup := app.Group("/v1/user")
	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)
}
