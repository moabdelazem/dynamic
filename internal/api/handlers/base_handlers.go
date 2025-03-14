package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// BaseHandler handles non-product HTTP requests
type BaseHandler struct{}

// NewBaseHandler creates a new base handler
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// RegisterRoutes registers the base routes
func (h *BaseHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.Home).Methods("GET")
	router.HandleFunc("/health", h.Health).Methods("GET")
	router.HandleFunc("/info", h.Info).Methods("GET")
}

// Home handles GET /
func (h *BaseHandler) Home(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}

// Health handles GET /health
func (h *BaseHandler) Health(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// Info handles GET /info
func (h *BaseHandler) Info(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"name":     "Mohamed Abdelazem",
		"email":    "mabdelazemahmed@gmail.com",
		"github":   "moabdelazem",
		"linkedin": "https://www.linkedin.com/in/moabdelazem/",
	})
}
