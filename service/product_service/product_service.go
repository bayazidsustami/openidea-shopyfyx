package product_service

import (
	"context"
	product_model "openidea-shopyfyx/models/product"
	user_model "openidea-shopyfyx/models/user"
)

type ProductService interface {
	Create(ctx context.Context, user user_model.User, request product_model.CreateProductRequest) error
	Update(ctx context.Context, user user_model.User, request product_model.UpdateProductRequest) error
	Delete(ctx context.Context, user user_model.User, productId int) error
	GetAllProducts(ctx context.Context, user user_model.User, filterProduct product_model.FilterProducts) (product_model.PagingProductResponse, error)
	GetProductById(ctx context.Context, user user_model.User, productId int) (product_model.ProductUsersResponse, error)
	UpdateProductStock(ctx context.Context, user user_model.User, productId int, request product_model.UpdateProductStockRequest) error
	BuyProduct(ctx context.Context, user user_model.User, productId int, request product_model.ProductPaymentRequest) error
}
