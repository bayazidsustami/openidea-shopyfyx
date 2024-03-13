package product_model

type FilterProducts struct {
	UserOnly       bool      `json:"userOnly"`
	Limit          int       `json:"limit"`
	Offset         int       `json:"offset"`
	Tags           []string  `json:"tags"`
	Condition      Condition `json:"condition"`
	ShowEmptyStock bool      `json:"showEmptyStock"`
	MaxPrice       float64   `json:"maxPrice"`
	MinPrice       float64   `json:"minPrice"`
	SortBy         string    `json:"sortBy"`
	OrderBy        string    `json:"orderBy"`
	Search         string    `json:"search"`
}
