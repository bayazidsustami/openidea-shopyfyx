package product_controller

import (
	product_model "openidea-shopyfyx/models/product"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/product_service"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService product_service.ProductService
	AuthService    auth_service.AuthService
}

func New(
	productService product_service.ProductService,
	authService auth_service.AuthService,
) ProductController {
	return ProductController{
		ProductService: productService,
		AuthService:    authService,
	}
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	productRequest := new(product_model.CreateProductRequest)

	err := ctx.BodyParser(productRequest)
	if err != nil {
		return err
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return err
	}

	err = controller.ProductService.Create(ctx.UserContext(), user, *productRequest)
	if err != nil {
		return err
	}

	return ctx.SendString("success")
}
