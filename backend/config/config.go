package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	ServerPort     int
	GitHubToken    string
	OpenAIKey      string
	DatabaseURL    string
	LogLevel       string
	MaxReportCount int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // default port
	}

	maxReports, err := strconv.Atoi(os.Getenv("MAX_REPORTS"))
	if err != nil {
		maxReports = 100 // default max reports
	}

	return &Config{
		ServerPort:     port,
		GitHubToken:    os.Getenv("GITHUB_TOKEN"),
		OpenAIKey:      os.Getenv("OPENAI_API_KEY"),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		LogLevel:       os.Getenv("LOG_LEVEL"),
		MaxReportCount: maxReports,
	}
}
