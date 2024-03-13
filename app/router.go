package app

import (
	"openidea-shopyfyx/controller/product_controller"
	"openidea-shopyfyx/controller/user_controller"
	"openidea-shopyfyx/db"
	product_repository "openidea-shopyfyx/repository/product"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/product_service"
	"openidea-shopyfyx/service/user_service"
	"openidea-shopyfyx/utils"

	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
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

	productRepository := product_repository.New()
	productService := product_service.New(dbPool, validator, productRepository)
	productController := product_controller.New(productService, authService)

	userGroup := app.Group("/v1/user")
	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)

	productRoute := app.Group("/v1/product", getJwtTokenHandler())
	productRoute.Get("/", func(c *fiber.Ctx) error { return err })
	productRoute.Post("/", productController.Create)
	productRoute.Patch("/:productId", productController.Update)
	productRoute.Delete("/:productId", func(c *fiber.Ctx) error { return err })

}

// TODO jangan lupa update secrets key
func getJwtTokenHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("ini rahasia")},
		ContextKey: "userInfo",
	})
}
