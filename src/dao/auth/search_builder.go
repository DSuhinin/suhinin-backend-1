package auth

import (
	"strconv"
	"strings"
)

// EventSearchBuilder represents search params.
type EventSearchBuilder struct {
	ID         string
	CustomerID string
	FromDate   string
	ToDate     string
	Type       string
	Country    string
	Region     string
	City       string
	PostCode   string
	Page       string
	Limit      string
}

// Build builds the `WHERE` part of SQL to make a servers search.
func (ss EventSearchBuilder) Build() (string, []interface{}, int, int) {
	statements, params := []string{"1=1"}, []interface{}{}
	if ss.ID != "" {
		params = append(params, ss.ID)
		statements = append(statements, "e.id = ?")
	}
	if ss.Type != "" {
		params = append(params, ss.Type)
		statements = append(statements, "e.type = ?")
	}
	if ss.FromDate != "" {
		params = append(params, ss.FromDate)
		statements = append(statements, "e.created_at >= ?")
	}
	if ss.ToDate != "" {
		params = append(params, ss.ToDate)
		statements = append(statements, "e.created_at <= ?")
	}
	if ss.CustomerID != "" {
		params = append(params, ss.CustomerID)
		statements = append(statements, "s.customer_id = ?")
	}
	if ss.Country != "" {
		params = append(params, ss.Country)
		statements = append(statements, "s.country = ?")
	}
	if ss.Region != "" {
		params = append(params, ss.Region)
		statements = append(statements, "s.region = ?")
	}
	if ss.City != "" {
		params = append(params, ss.City)
		statements = append(statements, "s.city = ?")
	}
	if ss.PostCode != "" {
		params = append(params, ss.PostCode)
		statements = append(statements, "s.post_code = ?")
	}

	limit, offset := 10, 0
	if ss.Limit != "" {
		number, err := strconv.Atoi(ss.Limit)
		if err == nil {
			limit = number
		}
	}
	if ss.Page != "" {
		number, err := strconv.Atoi(ss.Page)
		if err == nil {
			offset = (number - 1) * limit
		}
	}

	return strings.Join(statements, " AND "), params, limit, offset
}
