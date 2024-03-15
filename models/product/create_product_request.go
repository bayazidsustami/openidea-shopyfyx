package product_model

type CreateProductRequest struct {
	Name           string    `json:"name" validate:"required,min=5,max=60"`
	Price          float64   `json:"price" validate:"required,min=0"`
	ImageUrl       string    `json:"imageUrl" validate:"required,url"`
	Stock          int       `json:"stock" validate:"required,min=0"`
	Condition      Condition `json:"condition" validate:"required,oneof=new second"`
	Tags           []string  `json:"tags" validate:"required,min=0"`
	IsPurchaseable *bool     `json:"isPurchaseable" validate:"required,boolean"`
}
