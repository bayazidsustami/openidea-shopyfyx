package product_repository

import (
	"context"
	product_model "openidea-shopyfyx/models/product"

	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	Create(ctx context.Context, tx pgx.Tx, product product_model.Product) (product_model.Product, error)
	Update(ctx context.Context, tx pgx.Tx, product product_model.Product) (product_model.Product, error)
	Delete(ctx context.Context, tx pgx.Tx, userId int, productId int) error
	GetAllProduct(ctx context.Context, tx pgx.Tx, userId int, filterProduct product_model.FilterProducts) ([]product_model.Product, product_model.MetaPage, error)
	GetProductById(ctx context.Context, tx pgx.Tx, productId int) (product_model.ProductUsers, error)
	UpdateProductStock(ctx context.Context, tx pgx.Tx, userId int, productId int, stockAmount int) error
	BuyProduct(ctx context.Context, tx pgx.Tx, userId int, productId int, request product_model.ProductPaymentRequest) error
}
