package app

import (
	"openidea-shopyfyx/controller/bank_account_controller"
	"openidea-shopyfyx/controller/image_upload_controller"
	"openidea-shopyfyx/controller/product_controller"
	"openidea-shopyfyx/controller/user_controller"
	"openidea-shopyfyx/db"
	bank_account_repository "openidea-shopyfyx/repository/bank_account"
	product_repository "openidea-shopyfyx/repository/product"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/bank_account_service"
	"openidea-shopyfyx/service/image_service"
	"openidea-shopyfyx/service/product_service"
	"openidea-shopyfyx/service/user_service"
	"openidea-shopyfyx/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-playground/validator/v10"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func RegisterRoute(app *fiber.App) {

	validator := validator.New()
	registerValidation(validator)

	dbPool, err := db.InitDBPool()
	utils.PanicErr(err)

	authService := auth_service.New()

	userRepository := user_repository.New()
	userService := user_service.New(userRepository, validator, authService, dbPool)
	userController := user_controller.New(userService)

	productRepository := product_repository.New()
	productService := product_service.New(dbPool, validator, productRepository, userRepository)
	productController := product_controller.New(productService, authService)

	imageService := image_service.New(getAwsSession())
	imageUploadController := image_upload_controller.New(authService, imageService)

	bankAccountRepository := bank_account_repository.New(dbPool)
	bankAccountService := bank_account_service.New(bankAccountRepository, validator)
	bankAccountController := bank_account_controller.New(bankAccountService, authService)

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
	productRoute.Post("/:productId/stock", productController.UpdateProductStock)
	productRoute.Post("/:productId/buy", productController.BuyProduct)

	imageRoute := app.Group("/v1/image")
	imageRoute.Post("/", imageUploadController.UploadImage)

	bankAccountRoute := app.Group("/v1/bank/account")
	bankAccountRoute.Get("", bankAccountController.GetAllByUserId)
	bankAccountRoute.Post("", bankAccountController.Create)
	bankAccountRoute.Patch("/:bankAccountId", bankAccountController.Update)
	bankAccountRoute.Delete("/:bankAccountId", bankAccountController.Delete)
}

// TODO jangan lupa update secrets key
func getJwtTokenHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("JWT_SECRET"))},
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

func getAwsSession() *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			viper.GetString("S3_ID"),
			viper.GetString("S3_SECRET_KEY"),
			"",
		),
	}))
	svc := s3.New(sess)

	return svc
}

func registerValidation(validation *validator.Validate) {
	validation.RegisterValidation("imageurl", utils.MustValidImageUrl)
}
