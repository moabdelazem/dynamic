package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/moabdelazem/dynamicdevops/internal/api"
	"github.com/moabdelazem/dynamicdevops/pkg/config"
)

func main() {
	// Configure logger
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("INFO: Starting Dynamic DevOps API server")

	// Load configuration
	cfg := config.NewConfig()

	// Setup router
	router := api.SetupRouter()

	// Log server startup
	log.Printf("INFO: Server listening on port %d", cfg.Port)

	// Start server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Printf("ERROR: Failed to start server: %v", err)
		os.Exit(1)
	}
}
