package product_repository

import (
	"context"
	product_model "openidea-shopyfyx/models/product"

	"github.com/jackc/pgx/v5"
)

type ProductRepositoryImpl struct {
}

func New() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, product product_model.Product) product_model.Product {

}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, product product_model.Product) product_model.Product {
	return product_model.Product{}
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, productId int) {

}
func (repository *ProductRepositoryImpl) GetAllProduct(ctx context.Context, tx pgx.Tx) []product_model.Product {
	return []product_model.Product{}
}

func (repository *ProductRepositoryImpl) GetProductById(ctx context.Context, tx pgx.Tx, productId int) product_model.Product {
	return product_model.Product{}
}
