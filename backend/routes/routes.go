package routes

import (
	"hackathon-ai-auditor-backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *handlers.Handlers) {
	// Health check endpoint
	router.GET("/health", h.HealthCheck)

	// Code analysis endpoint
	router.POST("/analyze", h.AnalyzeCode)

	// GitHub webhook endpoint
	router.POST("/github/webhook", h.GitHubWebhook)

	// Additional endpoints for the AI Code Auditor
	router.POST("/analyze/batch", h.BatchAnalysis)
	router.GET("/reports/:id", h.GetReport)
	router.GET("/reports", h.ListReports)
}
