package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moabdelazem/dynamicdevops/internal/models"
)

// RespondWithJSON sends a JSON response with the given status code
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("ERROR: Failed to encode JSON response: %v", err)
	}
}

// RespondWithError sends a JSON error response
func RespondWithError(w http.ResponseWriter, statusCode int, message string, err error) {
	apiError := models.APIError{
		Status:  statusCode,
		Message: message,
	}
	if err != nil {
		apiError.Error = err.Error()
	}
	RespondWithJSON(w, statusCode, apiError)
}
