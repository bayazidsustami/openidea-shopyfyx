package product_controller

import (
	product_model "openidea-shopyfyx/models/product"
	"openidea-shopyfyx/service/auth_service"
	"openidea-shopyfyx/service/product_service"
	"strconv"

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

func (controller *ProductController) Update(ctx *fiber.Ctx) error {
	productRequest := new(product_model.UpdateProductRequest)
	productIdString := ctx.Params("productId")

	productId, err := strconv.Atoi(productIdString)
	if err != nil {
		return err
	}

	productRequest.ProductId = productId

	err = ctx.BodyParser(productRequest)
	if err != nil {
		return err
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "something error")
	}

	err = controller.ProductService.Update(ctx.UserContext(), user, *productRequest)
	if err != nil {
		return err
	}

	return ctx.SendString("success")
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	productIdString := ctx.Params("productId")

	productId, err := strconv.Atoi(productIdString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "something error")
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "something error")
	}

	err = controller.ProductService.Delete(ctx.UserContext(), user, productId)
	if err != nil {
		return err
	}

	return ctx.SendString("success")
}

func (controller *ProductController) GetAllProducts(ctx *fiber.Ctx) error {

	filterProduct := new(product_model.FilterProducts)

	err := ctx.QueryParser(filterProduct)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "something error")
	}
	products, err := controller.ProductService.GetAllProducts(ctx.UserContext(), user, *filterProduct)
	if err != nil {
		return err
	}

	return ctx.JSON(products)
}

func (controller *ProductController) GetProductById(ctx *fiber.Ctx) error {
	productIdString := ctx.Params("productId")

	productId, err := strconv.Atoi(productIdString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "something error")
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "something error")
	}

	result, err := controller.ProductService.GetProductById(ctx.UserContext(), user, productId)
	if err != nil {
		return err
	}

	return ctx.JSON(result)
}

func (controller *ProductController) UpdateProductStock(ctx *fiber.Ctx) error {
	updateProductStock := new(product_model.UpdateProductStockRequest)

	err := ctx.BodyParser(updateProductStock)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	productIdString := ctx.Params("productId")
	productId, err := strconv.Atoi(productIdString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "something error")
	}

	user, err := controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "something error")
	}

	err = controller.ProductService.UpdateProductStock(ctx.UserContext(), user, productId, *updateProductStock)
	if err != nil {
		return err
	}

	return ctx.SendString("success")
}

func (controller *ProductController) BuyProduct(ctx *fiber.Ctx) error {
	paymentRequest := new(product_model.ProductPaymentRequest)

	err := ctx.BodyParser(paymentRequest)

	productIdString := ctx.Params("productId")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	_, err = strconv.Atoi(productIdString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "something error")
	}

	_, err = controller.AuthService.GetValidUser(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "something error")
	}

	return ctx.SendString("oke")
}
