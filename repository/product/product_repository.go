package product_repository

import (
	"context"
	product_model "openidea-shopyfyx/models/product"

	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	Create(ctx context.Context, tx pgx.Tx, product product_model.Product) product_model.Product
	Update(ctx context.Context, tx pgx.Tx, product product_model.Product) product_model.Product
	Delete(ctx context.Context, tx pgx.Tx, userId int, productId int)
	GetAllProduct(ctx context.Context, tx pgx.Tx, userId int) []product_model.Product
	GetProductById(ctx context.Context, tx pgx.Tx, userId int, productId int) product_model.Product
}
