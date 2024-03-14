package product_repository

import (
	"context"
	product_model "openidea-shopyfyx/models/product"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type ProductRepositoryImpl struct {
}

func New() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, product product_model.Product) (product_model.Product, error) {
	PROUDCT_INSERT := "INSERT INTO products(product_name, condition, price, tags, is_available, image_url, user_id)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING product_id"

	PRODUCT_STOCK_INSERT := "INSERT INTO product_stocks(product_id, quantity) VALUES($1, $2)"

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
	if err != nil {
		return product_model.Product{}, err
	}

	_, err = tx.Exec(ctx, PRODUCT_STOCK_INSERT, productId, product.ProductStock.Quantity)
	if err != nil {
		return product_model.Product{}, err
	}

	product.ProductId = productId

	return product, nil
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, product product_model.Product) (product_model.Product, error) {
	PRODUCT_UPDATE := "UPDATE products " +
		"SET product_name = $1,  price = $2, image_url = $3, tags = $4, is_available=$5, updated_at = CURRENT_TIMESTAMP " +
		"WHERE product_id = $6 AND user_id = $7"

	productTags := strings.Join(product.Tags, ",")

	result, err := tx.Exec(ctx, PRODUCT_UPDATE,
		product.ProductName,
		product.Price,
		product.ImageUrl,
		productTags,
		product.IsAvailable,
		product.ProductId,
		product.UserId,
	)

	if err != nil {
		return product_model.Product{}, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return product_model.Product{}, fiber.NewError(fiber.StatusNotFound, "not found id : "+strconv.Itoa(product.ProductId))
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, userId int, productId int) error {
	PRODUCT_DELETE := "UPDATE products SET deleted_at = CURRENT_TIMESTAMP WHERE product_id=$1"
	result, err := tx.Exec(ctx, PRODUCT_DELETE, productId)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if result.RowsAffected() == 0 {
		return fiber.NewError(fiber.StatusNotFound, "not found id : "+strconv.Itoa(productId))
	}

	return nil
}

func (repository *ProductRepositoryImpl) GetAllProduct(ctx context.Context, tx pgx.Tx, userId int, filterProduct product_model.FilterProducts) ([]product_model.Product, product_model.MetaPage, error) {
	query := filterProduct.BuildQuery(userId)

	rows, err := tx.Query(ctx, query, userId)
	if err != nil {
		return nil, product_model.MetaPage{}, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}
	defer rows.Close()

	var products []product_model.Product
	for rows.Next() {
		var tags string
		product := product_model.Product{}
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.Condition,
			&tags,
			&product.IsAvailable,
			&product.ImageUrl,
			&product.UserId,
			&product.ProductStock.ProductId,
			&product.ProductStock.Quantity,
		)
		product.Tags = strings.Split(tags, ",")
		if err != nil {
			return nil, product_model.MetaPage{}, fiber.NewError(fiber.StatusInternalServerError, "something error")
		}
		products = append(products, product)
	}
	return products, product_model.MetaPage{
		Limit:  filterProduct.Limit,
		Offset: filterProduct.Offset,
		Total:  len(products),
	}, nil
}

func (repository *ProductRepositoryImpl) GetProductById(ctx context.Context, tx pgx.Tx, userId int, productId int) (product_model.Product, error) {
	GET_PRODUCT := "SELECT p.product_id, p.product_name, p.price, p.condition, p.tags, p.is_available, p.image_url, p.user_id, ps.product_stock_id, ps.quantity " +
		"FROM products p " +
		"JOIN product_stocks ps ON p.product_id = ps.product_id " +
		"WHERE p.deleted_at IS NULL " +
		"AND p.product_id = $1" +
		"AND p.user_id = $2"
	product := product_model.Product{}
	err := tx.QueryRow(ctx, GET_PRODUCT, productId, userId).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.Price,
		&product.Condition,
		&product.Tags,
		&product.IsAvailable,
		&product.ImageUrl,
		&product.UserId,
		&product.ProductStock.ProductId,
		&product.ProductStock.Quantity,
	)
	if err != nil {
		return product_model.Product{}, fiber.NewError(fiber.StatusInternalServerError, "something error")
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) UpdateProductStock(ctx context.Context, tx pgx.Tx, userId int, productId int, stockAmount int) error {
	UPDATE_PRODUCT_STOCK := "UPDATE product_stocks AS ps " +
		"SET updated_at = CURRENT_TIMESTAMP, quantity = $1 " +
		"FROM products AS p" +
		"WHERE ps.product_id = $2" +
		"AND p.user_id = $3"

	result, err := tx.Exec(ctx, UPDATE_PRODUCT_STOCK, stockAmount, productId, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if result.RowsAffected() == 0 {
		return fiber.NewError(fiber.StatusNotFound, "not found id : "+strconv.Itoa(productId))
	}

	return nil
}
