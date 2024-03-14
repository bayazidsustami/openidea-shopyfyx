package product_model

type UpdateProductStockRequest struct {
	Stock int `json:"stock" validate:"required,min=0"`
}
