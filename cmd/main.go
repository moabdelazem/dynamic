package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PORT = 8080
)

// Logger middleware to log HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request details after handling
		log.Printf("REQUEST: %s %s - %s - %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

func main() {
	// Configure logger
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("INFO: Starting Dynamic DevOps API server")

	router := http.NewServeMux()

	// Register routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Send JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "Hello, World!"}`)
	})

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	// Info Route
	// Return Some Info About Me
	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"name": "Mohamed Abdelazem", "email": "mabdelazemahmed@gmail.com", "github": "moabdelazem", "linkedin": "https://www.linkedin.com/in/moabdelazem/"}`)
	})

	// Apply middleware
	loggedRouter := loggingMiddleware(router)

	// Log server startup
	log.Printf("INFO: Server listening on port %d", PORT)

	// Start server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), loggedRouter); err != nil {
		log.Printf("ERROR: Failed to start server: %v", err)
		os.Exit(1)
	}
}
