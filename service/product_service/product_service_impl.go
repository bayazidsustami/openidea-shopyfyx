package product_service

import (
	"context"
	product_model "openidea-shopyfyx/models/product"
	user_model "openidea-shopyfyx/models/user"
	product_repository "openidea-shopyfyx/repository/product"
	"openidea-shopyfyx/utils"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductServiceImpl struct {
	DBPool            *pgxpool.Pool
	Validator         *validator.Validate
	ProductRepository product_repository.ProductRepository
}

func New(
	DBPool *pgxpool.Pool,
	validator *validator.Validate,
	productRepository product_repository.ProductRepository,
) ProductService {
	return &ProductServiceImpl{
		DBPool:            DBPool,
		Validator:         validator,
		ProductRepository: productRepository,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, user user_model.User, request product_model.CreateProductRequest) error {
	err := service.Validator.Struct(request)
	if err != nil {
		return err
	}

	conn, err := service.DBPool.Acquire(ctx)
	utils.PanicErr(err)
	defer conn.Release()

	tx, err := conn.Begin(ctx)
	utils.PanicErr(err)
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
		IsAvailable: request.IsPurchaseable,
	}

	_ = service.ProductRepository.Create(ctx, tx, product)
	return nil
}
