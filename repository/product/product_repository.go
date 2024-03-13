package product_repository

import (
	"context"
	"openidea-shopyfyx/models"
	product_model "openidea-shopyfyx/models/product"

	"github.com/jackc/pgx/v5"
)

type ProductRepository interface {
	Create(ctx context.Context, tx pgx.Tx, product product_model.Product) (product_model.Product, error)
	Update(ctx context.Context, tx pgx.Tx, product product_model.Product) (product_model.Product, error)
	Delete(ctx context.Context, tx pgx.Tx, userId int, productId int) error
	GetAllProduct(ctx context.Context, tx pgx.Tx, userId int, pageInfo models.MetaPageRequest) ([]product_model.Product, error)
	GetProductById(ctx context.Context, tx pgx.Tx, userId int, productId int) (product_model.Product, error)
}
