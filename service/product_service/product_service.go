package product_service

import (
	"context"
	product_model "openidea-shopyfyx/models/product"
	user_model "openidea-shopyfyx/models/user"
)

type ProductService interface {
	Create(ctx context.Context, user user_model.User, request product_model.CreateProductRequest) error
}
