package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/moabdelazem/dynamicdevops/internal/models"
	"github.com/moabdelazem/dynamicdevops/internal/store"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	store *store.ProductStore
}

// NewProductHandler creates a new product handler
func NewProductHandler(store *store.ProductStore) *ProductHandler {
	return &ProductHandler{
		store: store,
	}
}

// RegisterRoutes registers the product routes
func (h *ProductHandler) RegisterRoutes(router *mux.Router) {
	productRouter := router.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", h.ListProducts).Methods("GET")
	productRouter.HandleFunc("", h.CreateProduct).Methods("POST")
	productRouter.HandleFunc("/{id}", h.GetProduct).Methods("GET")
	productRouter.HandleFunc("/{id}", h.UpdateProduct).Methods("PUT")
	productRouter.HandleFunc("/{id}", h.DeleteProduct).Methods("DELETE")
}

// ListProducts handles GET /products
func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page := 1
	limit := 10

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if pageVal, err := strconv.Atoi(pageStr); err == nil && pageVal > 0 {
			page = pageVal
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limitVal, err := strconv.Atoi(limitStr); err == nil && limitVal > 0 {
			limit = limitVal
		}
	}

	products := h.store.ListProducts(page, limit)

	// Create paginated response
	response := models.PaginatedResponse{
		Data: products,
	}
	response.Pagination.Page = page
	response.Pagination.Limit = limit
	response.Pagination.TotalRows = h.store.GetTotalProducts()
	response.Pagination.TotalPage = (h.store.GetTotalProducts() + limit - 1) / limit

	RespondWithJSON(w, http.StatusOK, response)
}

// GetProduct handles GET /products/{id}
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, ok := h.store.GetProduct(id)
	if !ok {
		RespondWithError(w, http.StatusNotFound, "Product not found", nil)
		return
	}
	RespondWithJSON(w, http.StatusOK, product)
}

// CreateProduct handles POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.store.CreateProduct(product); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to create product", err)
		return
	}

	RespondWithJSON(w, http.StatusCreated, product)
}

// UpdateProduct handles PUT /products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updates models.Product
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := h.store.UpdateProduct(id, updates); err != nil {
		if strings.Contains(err.Error(), "not found") {
			RespondWithError(w, http.StatusNotFound, "Product not found", nil)
		} else {
			RespondWithError(w, http.StatusBadRequest, "Failed to update product", err)
		}
		return
	}

	product, _ := h.store.GetProduct(id)
	RespondWithJSON(w, http.StatusOK, product)
}

// DeleteProduct handles DELETE /products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.store.DeleteProduct(id); err != nil {
		RespondWithError(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
