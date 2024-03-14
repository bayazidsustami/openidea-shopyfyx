package product_model

type FilterProducts struct {
	UserOnly       bool      `json:"userOnly" validate:"boolean"`
	Limit          int       `json:"limit" validate:"number"`
	Offset         int       `json:"offset" validate:"number"`
	Tags           []string  `json:"tags" validate:"dive"`
	Condition      Condition `json:"condition" validate:"oneof=new second ''"`
	ShowEmptyStock bool      `json:"showEmptyStock" validate:"boolean"`
	MaxPrice       float64   `json:"maxPrice" validate:"number"`
	MinPrice       float64   `json:"minPrice" validate:"number"`
	SortBy         string    `json:"sortBy" validate:"oneof=price date ''"`
	OrderBy        string    `json:"orderBy" validate:"oneof=asc dsc ''"`
	Search         string    `json:"search"`
}
