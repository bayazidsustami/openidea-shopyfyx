package product_service

import (
	"context"
	"openidea-shopyfyx/models"
	product_model "openidea-shopyfyx/models/product"
	user_model "openidea-shopyfyx/models/user"
)

type ProductService interface {
	Create(ctx context.Context, user user_model.User, request product_model.CreateProductRequest) error
	Update(ctx context.Context, user user_model.User, request product_model.UpdateProductRequest) error
	Delete(ctx context.Context, user user_model.User, productId int) error
	GetAllProducts(ctx context.Context, user user_model.User, pageInfo models.MetaPageRequest) (product_model.PagingProductResponse, error)
	GetProductById(ctx context.Context, user user_model.User, productId int) error
}
