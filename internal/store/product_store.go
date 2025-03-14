package store

import (
	"fmt"
	"time"

	"github.com/moabdelazem/dynamicdevops/internal/models"
)

// Products map to store products by ID
type Products map[string]models.Product

// ProductStore manages the product data
type ProductStore struct {
	products Products
}

// NewProductStore creates a new product store
func NewProductStore() *ProductStore {
	return &ProductStore{
		products: make(Products),
	}
}

// CreateProduct adds a new product to the store
func (s *ProductStore) CreateProduct(product models.Product) error {
	if product.ID == "" {
		return fmt.Errorf("product ID cannot be empty")
	}
	if product.Name == "" {
		return fmt.Errorf("product name cannot be empty")
	}
	if product.Price < 0 {
		return fmt.Errorf("product price cannot be negative")
	}

	// Set timestamps
	now := time.Now()
	if product.CreatedAt.IsZero() {
		product.CreatedAt = now
	}
	product.UpdatedAt = now

	s.products[product.ID] = product
	return nil
}

// UpdateProduct updates an existing product
func (s *ProductStore) UpdateProduct(id string, updates models.Product) error {
	product, exists := s.products[id]
	if !exists {
		return fmt.Errorf("product not found")
	}

	// Update fields if provided
	if updates.Name != "" {
		product.Name = updates.Name
	}
	if updates.Description != "" {
		product.Description = updates.Description
	}
	if updates.Price > 0 {
		product.Price = updates.Price
	}

	product.UpdatedAt = time.Now()
	s.products[id] = product
	return nil
}

// DeleteProduct removes a product from the store
func (s *ProductStore) DeleteProduct(id string) error {
	if _, exists := s.products[id]; !exists {
		return fmt.Errorf("product not found")
	}
	delete(s.products, id)
	return nil
}

// GetProduct retrieves a product by ID
func (s *ProductStore) GetProduct(id string) (models.Product, bool) {
	product, ok := s.products[id]
	return product, ok
}

// ListProducts returns all products with optional pagination
func (s *ProductStore) ListProducts(page, limit int) []models.Product {
	products := make([]models.Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}

	// Apply pagination if requested
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit

		// Bounds checking
		if start >= len(products) {
			return []models.Product{}
		}
		if end > len(products) {
			end = len(products)
		}

		return products[start:end]
	}

	return products
}

// GetTotalProducts returns the total number of products
func (s *ProductStore) GetTotalProducts() int {
	return len(s.products)
}
