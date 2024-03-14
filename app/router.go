package app

import (
	"openidea-shopyfyx/controller/image_upload_controller"
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

	userRepository := user_repository.New()
	userService := user_service.New(userRepository, validator, authService, dbPool)
	userController := user_controller.New(userService)

	productRepository := product_repository.New()
	productService := product_service.New(dbPool, validator, productRepository)
	productController := product_controller.New(productService, authService)

	imageUploadController := image_upload_controller.New(authService)

	userGroup := app.Group("/v1/user")
	userGroup.Post("/register", userController.Register)
	userGroup.Post("/login", userController.Login)

	app.Use(checkTokenHeaderExist)
	app.Use(getJwtTokenHandler())

	productRoute := app.Group("/v1/product")
	productRoute.Get("/", productController.GetAllProducts)
	productRoute.Get("/:productId", productController.GetProductById)
	productRoute.Post("/", productController.Create)
	productRoute.Patch("/:productId", productController.Update)
	productRoute.Delete("/:productId", productController.Delete)

	imageRoute := app.Group("/v1/image")
	imageRoute.Post("/", imageUploadController.UploadImage)

}

// TODO jangan lupa update secrets key
func getJwtTokenHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("ini rahasia")},
		ContextKey: "userInfo",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		},
	})
}

func checkTokenHeaderExist(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	} else {
		return ctx.Next()
	}
}
