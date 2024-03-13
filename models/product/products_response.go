package product_model

type ProductResponse struct {
	ProductId      int       `json:"productId"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	ImageUrl       string    `json:"imageUrl"`
	Stock          int       `json:"stock"`
	Condition      Condition `json:"condition"`
	Tags           []string  `json:"tags"`
	IsPurchaseable bool      `json:"isPurchaseable"`
	PurchaseCount  int       `json:"purchaseCount"`
}

type PagingProductResponse struct {
	Message  string            `json:"message"`
	Data     []ProductResponse `json:"data"`
	MetaPage MetaPage          `json:"meta"`
}

type MetaPage struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
