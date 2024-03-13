package product_repository

import (
	"context"
	product_model "openidea-shopyfyx/models/product"
	"openidea-shopyfyx/utils"
	"strings"

	"github.com/jackc/pgx/v5"
)

type ProductRepositoryImpl struct {
}

func New() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, product product_model.Product) product_model.Product {
	PROUDCT_INSERT := "INSERT INTO proudcts(product_name, condition, price, tags, is_available, image_url, user_id)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8) WHERE user_id = $9 RETURNING product_id"

	PRODUCT_STOCK_INSERT := "INSERT INTO prouduct_stocks(product_id, quantity) VALUES($1, $2)"

	productTags := strings.Join(product.Tags, ",")

	var productId int
	err := tx.QueryRow(ctx, PROUDCT_INSERT,
		product.ProductName,
		product.Condition,
		product.Price,
		productTags,
		product.IsAvailable,
		product.ImageUrl,
		product.UserId,
	).Scan(&productId)
	utils.PanicErr(err)

	_, err = tx.Exec(ctx, PRODUCT_STOCK_INSERT, product.ProductId, product.ProductStock.Quantity)
	utils.PanicErr(err)

	product.ProductId = productId

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, product product_model.Product) product_model.Product {
	PRODUCT_UPDATE := "UPDATE products " +
		"SET product_name = $1,  price = $2, image_url = $3, tags = $4, is_available=$5, update_at = CURRENT_TIMESTAMP" +
		"WHERE product_id = $6 AND user_id = $7"

	productTags := strings.Join(product.Tags, ",")

	_, err := tx.Exec(ctx, PRODUCT_UPDATE,
		product.ProductName,
		product.Price,
		product.ImageUrl,
		productTags,
		product.IsAvailable,
		product.ProductId,
		product.UserId,
	)
	utils.PanicErr(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, productId int) {
	PRODUCT_DELETE := "UPDATE products SET delete_at = CURRENT_TIMESTAMP WHERE product_id=$1"
	_, err := tx.Exec(ctx, PRODUCT_DELETE, productId)
	utils.PanicErr(err)
}
func (repository *ProductRepositoryImpl) GetAllProduct(ctx context.Context, tx pgx.Tx) []product_model.Product {
	return []product_model.Product{}
}

func (repository *ProductRepositoryImpl) GetProductById(ctx context.Context, tx pgx.Tx, productId int) product_model.Product {
	return product_model.Product{}
}
