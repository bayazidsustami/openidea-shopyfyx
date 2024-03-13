package product_model

import "time"

type Product struct {
	ProductId    int
	ProductName  string
	Condition    Condition
	Tags         []string
	IsAvailable  bool
	ImageUrl     string
	UserId       int
	CreatedAt    *time.Time
	DeletedAt    *time.Time
	UpdatedAt    *time.Time
	ProductStock ProductStock
}

type Condition string

const (
	New    Condition = "new"
	Second Condition = "second"
)
