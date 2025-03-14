package api

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/moabdelazem/dynamicdevops/internal/api/handlers"
	"github.com/moabdelazem/dynamicdevops/internal/api/middleware"
	"github.com/moabdelazem/dynamicdevops/internal/models"
	"github.com/moabdelazem/dynamicdevops/internal/store"
)

// SetupRouter sets up the router with all routes and middleware
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Create stores
	productStore := store.NewProductStore()

	// Add sample products
	productStore.CreateProduct(models.Product{
		ID:          "1",
		Name:        "Product 1",
		Description: "Product 1 Description",
		Price:       100,
		CreatedAt:   time.Now(),
	})

	productStore.CreateProduct(models.Product{
		ID:          "2",
		Name:        "Product 2",
		Description: "Product 2 Description",
		Price:       200,
		CreatedAt:   time.Now(),
	})

	productStore.CreateProduct(models.Product{
		ID:          "3",
		Name:        "Product 3",
		Description: "Product 3 Description",
		Price:       300,
		CreatedAt:   time.Now(),
	})

	// Create handlers
	baseHandler := handlers.NewBaseHandler()
	productHandler := handlers.NewProductHandler(productStore)

	// Register routes
	baseHandler.RegisterRoutes(router)
	productHandler.RegisterRoutes(router)

	// Apply middleware
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)

	return router
}
