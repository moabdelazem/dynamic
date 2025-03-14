package models

import (
	"time"
)

// Product represents a product in our store
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// PaginatedResponse wraps a list response with pagination metadata
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination struct {
		Page      int `json:"page"`
		Limit     int `json:"limit"`
		TotalRows int `json:"total_rows"`
		TotalPage int `json:"total_pages"`
	} `json:"pagination"`
}

// APIError represents a structured error response
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
