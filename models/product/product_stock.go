package product_model

import "time"

type ProductStock struct {
	ProductStockId int
	ProductId      int
	Quantity       int
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
