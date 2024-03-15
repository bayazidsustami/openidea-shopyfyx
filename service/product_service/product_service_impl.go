package product_service

import (
	"context"
	bank_account_model "openidea-shopyfyx/models/bank_account"
	product_model "openidea-shopyfyx/models/product"
	user_model "openidea-shopyfyx/models/user"
	product_repository "openidea-shopyfyx/repository/product"
	user_repository "openidea-shopyfyx/repository/user"
	"openidea-shopyfyx/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductServiceImpl struct {
	DBPool            *pgxpool.Pool
	Validator         *validator.Validate
	ProductRepository product_repository.ProductRepository
	UserRepository    user_repository.UserRepository
}

func New(
	DBPool *pgxpool.Pool,
	validator *validator.Validate,
	productRepository product_repository.ProductRepository,
	userRepository user_repository.UserRepository,
) ProductService {
	return &ProductServiceImpl{
		DBPool:            DBPool,
		Validator:         validator,
		ProductRepository: productRepository,
		UserRepository:    userRepository,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, user user_model.User, request product_model.CreateProductRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(ctx, tx)

	product := product_model.Product{
		ProductName: request.Name,
		Price:       request.Price,
		Condition:   request.Condition,
		ImageUrl:    request.ImageUrl,
		ProductStock: product_model.ProductStock{
			Quantity: request.Stock,
		},
		UserId:      user.UserId,
		Tags:        request.Tags,
		IsAvailable: *request.IsPurchaseable,
	}

	_, err = service.ProductRepository.Create(ctx, tx, product)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, user user_model.User, request product_model.UpdateProductRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer utils.CommitOrRollback(ctx, tx)

	product := product_model.Product{
		ProductId:   request.ProductId,
		ProductName: request.Name,
		Price:       request.Price,
		Condition:   request.Condition,
		ImageUrl:    request.ImageUrl,
		UserId:      user.UserId,
		Tags:        request.Tags,
		IsAvailable: request.IsPurchaseable,
	}

	_, err = service.ProductRepository.Update(ctx, tx, product)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, user user_model.User, productId int) error {
	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer utils.CommitOrRollback(ctx, tx)

	err = service.ProductRepository.Delete(ctx, tx, user.UserId, productId)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProductServiceImpl) GetAllProducts(ctx context.Context, user user_model.User, filterProduct product_model.FilterProducts) (product_model.PagingProductResponse, error) {
	if err := service.Validator.Struct(filterProduct); err != nil {
		return product_model.PagingProductResponse{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return product_model.PagingProductResponse{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return product_model.PagingProductResponse{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer utils.CommitOrRollback(ctx, tx)

	if filterProduct.Limit == 0 {
		filterProduct.Limit = 10
	}

	if filterProduct.Limit == 0 {
		filterProduct.Offset = 1
	}

	products, meta, err := service.ProductRepository.GetAllProduct(ctx, tx, user.UserId, filterProduct)
	if err != nil {
		return product_model.PagingProductResponse{}, err
	}

	var pagingData product_model.PagingProductResponse
	for _, product := range products {
		productResponse := product_model.ProductResponse{
			ProductId:      product.ProductId,
			Name:           product.ProductName,
			Price:          product.Price,
			ImageUrl:       product.ImageUrl,
			Stock:          product.ProductStock.Quantity,
			Condition:      product.Condition,
			Tags:           product.Tags,
			IsPurchaseable: product.IsAvailable,
			PurchaseCount:  product.PurchaseCount,
		}
		pagingData.Data = append(pagingData.Data, productResponse)
	}
	pagingData.Message = "ok"
	pagingData.MetaPage = product_model.MetaPage{
		Limit:  meta.Limit,
		Offset: meta.Offset,
		Total:  meta.Total,
	}

	return pagingData, nil
}

func (service *ProductServiceImpl) GetProductById(ctx context.Context, user user_model.User, productId int) (product_model.ProductUsersResponse, error) {
	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return product_model.ProductUsersResponse{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return product_model.ProductUsersResponse{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer utils.CommitOrRollback(ctx, tx)

	productUser, err := service.ProductRepository.GetProductById(ctx, tx, productId)
	if err != nil {
		return product_model.ProductUsersResponse{}, err
	}

	seller, err := service.UserRepository.GetSeller(ctx, tx, productUser.Product.UserId)
	if err != nil {
		return product_model.ProductUsersResponse{}, err
	}

	productResponse := product_model.ProductUsersResponse{
		Product: product_model.ProductResponse{
			ProductId:      productUser.Product.ProductId,
			Name:           productUser.Product.ProductName,
			Price:          productUser.Product.Price,
			ImageUrl:       productUser.Product.ImageUrl,
			Stock:          productUser.Product.ProductStock.Quantity,
			Condition:      productUser.Product.Condition,
			Tags:           productUser.Product.Tags,
			IsPurchaseable: productUser.Product.IsAvailable,
			PurchaseCount:  productUser.Product.PurchaseCount,
		},
		Seller: product_model.Seller{
			Name:          productUser.Name,
			PurchaseTotal: seller.ProductsSoldTotal,
		},
	}

	for _, bank := range productUser.BankAccounts {
		bankAccount := bank_account_model.BankAccountData{
			BankAccountId:     bank.BankAccountId,
			BankName:          bank.BankName,
			BankAccountName:   bank.BankAccountName,
			BankAccountNumber: bank.BankAccountNumber,
		}
		productResponse.Seller.BankAccounts = append(productResponse.Seller.BankAccounts, bankAccount)
	}

	return productResponse, nil
}

func (service *ProductServiceImpl) UpdateProductStock(ctx context.Context, user user_model.User, productId int, request product_model.UpdateProductStockRequest) error {
	if err := service.Validator.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer utils.CommitOrRollback(ctx, tx)

	err = service.ProductRepository.UpdateProductStock(ctx, tx, user.UserId, productId, request.Stock)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProductServiceImpl) BuyProduct(ctx context.Context, user user_model.User, productId int, request product_model.ProductPaymentRequest) error {
	if err := service.Validator.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	conn, err := service.DBPool.Acquire(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer tx.Rollback(ctx)

	err = service.ProductRepository.BuyProduct(ctx, tx, user.UserId, productId, request)
	if err != nil {
		return err
	}

	return nil
}
