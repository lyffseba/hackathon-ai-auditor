package main

import (
	"fmt"
	"hackathon-ai-auditor-backend/analysis"
	"hackathon-ai-auditor-backend/config"
	"hackathon-ai-auditor-backend/handlers"
	"hackathon-ai-auditor-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create a Gin router with default middleware
	router := gin.Default()

	// Initialize analysis engine
	analyzer := analysis.NewAnalyzer(cfg.OpenAIKey)

	// Setup handlers with dependencies
	h := &handlers.Handlers{Analyzer: analyzer}

	// Setup routes with handlers
	routes.SetupRoutes(router, h)

	// Start the server with configured port
	router.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}
