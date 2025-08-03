package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hackathon-ai-auditor-backend/analysis"
	"hackathon-ai-auditor-backend/models"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v52/github"
)

// HealthCheck returns the status of the backend service
func (h *Handlers) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "AI Code Auditor backend is running",
		"timestamp": time.Now().UTC(),
	})
}

// Handlers contains all handler functions with dependencies
type Handlers struct {
	Analyzer      *analysis.Analyzer
	WebhookSecret string
}

// AnalyzeCode performs code analysis on the provided code snippet
func (h *Handlers) AnalyzeCode(c *gin.Context) {
	// Get the code from the request
	var request struct {
		Code     string `json:"code"`
		Language string `json:"language"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the analysis engine
	findings, err := h.Analyzer.AnalyzeCode(c.Request.Context(), request.Code, "input.code")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	report := models.Report{
		ID:        fmt.Sprintf("report_%d", time.Now().Unix()),
		Timestamp: time.Now().UTC(),
		Findings:  findings,
		Summary:   calculateSummary(findings),
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"report": report,
	})
}

// calculateSummary generates a summary from findings
func calculateSummary(findings []models.Finding) models.Summary {
	summary := models.Summary{
		TotalFindings: len(findings),
	}

	for _, f := range findings {
		switch f.Type {
		case "vulnerability", "security":
			summary.SecurityIssues++
		case "performance":
			summary.PerformanceIssues++
		case "best_practice", "code_smell":
			summary.BestPracticeViolations++
		}
	}

	return summary
}

// GitHubWebhook handles incoming GitHub webhooks
func (h *Handlers) GitHubWebhook(c *gin.Context) {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	eventType := c.GetHeader("X-GitHub-Event")
	signature := c.GetHeader("X-Hub-Signature-256")

	if err := h.VerifySignature(payload, signature); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}

	// Process the event based on eventType
	switch eventType {
	case "ping":
		c.JSON(http.StatusOK, gin.H{"status": "pong"})
	case "pull_request":
		var prEvent github.PullRequestEvent
		if err := json.Unmarshal(payload, &prEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}
		// Trigger analysis for the PR
		c.JSON(http.StatusOK, gin.H{"status": "success", "event": "pull_request"})
	default:
		c.JSON(http.StatusOK, gin.H{
			"status":     "success",
			"message":    "Webhook received",
			"event_type": eventType,
		})
	}
}

// BatchAnalysis performs analysis on multiple files or repositories
func (h *Handlers) BatchAnalysis(c *gin.Context) {
	// In a real implementation, this would handle batch analysis requests
	// For now, we'll return a mock report ID
	reportID := fmt.Sprintf("batch_report_%d", time.Now().Unix())

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"message":   "Batch analysis started",
		"report_id": reportID,
	})
}

// GetReport retrieves a specific analysis report by ID
func (h *Handlers) GetReport(c *gin.Context) {
	reportID := c.Param("id")

	// In a real implementation, this would retrieve a report from storage
	// For now, we'll return a mock report using our models
	findings := []models.Finding{
		{
			Type:     "security",
			Message:  "SQL injection vulnerability",
			Line:     15,
			Severity: "high",
			File:     "main.py",
		},
		{
			Type:     "performance",
			Message:  "Inefficient algorithm",
			Line:     23,
			Severity: "medium",
			File:     "utils.js",
		},
	}

	report := models.Report{
		ID:        reportID,
		Timestamp: time.Now().UTC(),
		Findings:  findings,
		Summary: models.Summary{
			TotalFindings:          len(findings),
			SecurityIssues:         1,
			PerformanceIssues:      1,
			BestPracticeViolations: 0,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"report": report,
	})
}

// ListReports returns a list of all analysis reports
func (h *Handlers) ListReports(c *gin.Context) {
	// In a real implementation, this would list reports from storage
	// For now, we'll return mock reports using our models

	reports := []models.Report{
		{
			ID:         "12345",
			Timestamp:  time.Now().Add(-5 * time.Minute).UTC(),
			Repository: "user/repo1",
			Findings: []models.Finding{
				{
					Type:     "security",
					Message:  "SQL injection vulnerability",
					Line:     15,
					Severity: "high",
					File:     "main.py",
				},
				{
					Type:     "performance",
					Message:  "Inefficient algorithm",
					Line:     23,
					Severity: "medium",
					File:     "utils.js",
				},
				{
					Type:     "best_practice",
					Message:  "Variable name could be more descriptive",
					Line:     8,
					Severity: "low",
					File:     "config.json",
				},
			},
			Summary: models.Summary{
				TotalFindings:          3,
				SecurityIssues:         1,
				PerformanceIssues:      1,
				BestPracticeViolations: 1,
			},
		},
		{
			ID:         "12346",
			Timestamp:  time.Now().Add(-10 * time.Minute).UTC(),
			Repository: "user/repo2",
			Findings: []models.Finding{
				{
					Type:     "security",
					Message:  "Hardcoded credentials detected",
					Line:     5,
					Severity: "critical",
					File:     "auth.js",
				},
				{
					Type:     "performance",
					Message:  "Memory leak in event handler",
					Line:     32,
					Severity: "high",
					File:     "app.js",
				},
				{
					Type:     "best_practice",
					Message:  "Missing error handling",
					Line:     45,
					Severity: "medium",
					File:     "api.py",
				},
				{
					Type:     "security",
					Message:  "XSS vulnerability in form input",
					Line:     18,
					Severity: "high",
					File:     "form.html",
				},
				{
					Type:     "performance",
					Message:  "Unnecessary database queries in loop",
					Line:     67,
					Severity: "medium",
					File:     "data.js",
				},
				{
					Type:     "best_practice",
					Message:  "Function is too long and complex",
					Line:     22,
					Severity: "low",
					File:     "main.go",
				},
				{
					Type:     "security",
					Message:  "Insecure random number generation",
					Line:     12,
					Severity: "medium",
					File:     "crypto.py",
				},
			},
			Summary: models.Summary{
				TotalFindings:          7,
				SecurityIssues:         3,
				PerformanceIssues:      2,
				BestPracticeViolations: 2,
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"reports": reports,
	})
}

// VerifySignature verifies the HMAC signature of the payload
func (h *Handlers) VerifySignature(payload []byte, signature string) error {
	if h.WebhookSecret == "" {
		return nil // Skip verification if no secret is set
	}

	hmac256 := hmac.New(sha256.New, []byte(h.WebhookSecret))
	hmac256.Write(payload)
	computed := hmac256.Sum(nil)

	signature = strings.TrimPrefix(signature, "sha256=")
	decoded, err := hex.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("invalid signature format")
	}

	if !hmac.Equal(decoded, computed) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
