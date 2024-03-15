package product_model

import (
	"fmt"
	"strings"
)

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

func (fp *FilterProducts) BuildQuery() string {
	query := "SELECT p.product_id, p.product_name, p.price, p.condition, p.tags, p.is_available, p.image_url, p.user_id, ps.product_stock_id, ps.quantity, " +
		"(SELECT COUNT(*) FROM orders o WHERE o.product_id = p.product_id) " +
		"FROM products p " +
		"JOIN product_stocks ps ON p.product_id = ps.product_id "

	// Initialize an empty slice to store the filter conditions
	conditions := []string{}

	if fp.UserOnly {
		conditions = append(conditions, "user_id = $1") //isUserOnly is param1
	}

	if fp.ShowEmptyStock {
		conditions = append(conditions, "ps.quantity = 0") //product
	}

	if fp.Tags != nil && len(fp.Tags) > 0 {
		tagsCondition := []string{}
		for _, tag := range fp.Tags {
			tagsCondition = append(tagsCondition, fmt.Sprintf("'%s' = ANY(string_to_array(p.tags, ','))", tag))
		}
		conditions = append(conditions, strings.Join(tagsCondition, " OR "))
	}

	if fp.Condition != "" {
		conditions = append(conditions, fmt.Sprintf("p.condition = '%s'", fp.Condition))
	}

	if fp.MaxPrice > 0 {
		conditions = append(conditions, fmt.Sprintf("p.price <= %.2f", fp.MaxPrice))
	}

	if fp.MinPrice > 0 {
		conditions = append(conditions, fmt.Sprintf("p.price >= %.2f", fp.MinPrice))
	}

	if fp.Search != "" {
		conditions = append(conditions, fmt.Sprintf("p.product_name LIKE '%%%s%%'", fp.Search))
	}

	conditions = append(conditions, "p.deleted_at IS NULL")

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Add sorting and ordering
	if fp.SortBy != "" {
		orderBy := "ASC"
		if fp.OrderBy == "dsc" {
			orderBy = "DESC"
		}
		mappedSortBy := "p.price"
		if fp.SortBy == "price" {
			mappedSortBy = "p.price"
		} else {
			mappedSortBy = "p.created_at"
		}
		query += fmt.Sprintf(" ORDER BY %s %s ", mappedSortBy, orderBy)
	}

	// Add limit and offset
	if fp.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", fp.Limit)
	}
	if fp.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", fp.Offset)
	}

	return query
}
