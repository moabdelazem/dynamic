package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Port int
}

// NewConfig creates a new configuration with values from environment variables
func NewConfig() *Config {
	return &Config{
		Port: getEnvAsInt("API_PORT", 8080),
	}
}

// Helper function to get an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Helper function to get an environment variable as an integer with a fallback value
func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
