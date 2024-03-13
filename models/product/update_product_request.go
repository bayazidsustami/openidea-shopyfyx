package product_model

type UpdateProductRequest struct {
	ProductId      int
	Name           string    `json:"name" validate:"required,min=5,max=60"`
	Price          float64   `json:"price" validate:"required,min=0"`
	ImageUrl       string    `json:"imageUrl" validate:"required,url"`
	Condition      Condition `json:"condition" validate:"required,oneof=new second"`
	Tags           []string  `json:"tags" validate:"required,min=0"`
	IsPurchaseable bool      `json:"isPurchaseable" validate:"required,boolean"`
}
