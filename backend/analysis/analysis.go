package analysis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"hackathon-ai-auditor-backend/models"

	"github.com/sashabaranov/go-openai"
)

// Analyzer is responsible for analyzing code using OpenAI API
type Analyzer struct {
	client *openai.Client
}

// NewAnalyzer creates a new instance of Analyzer
func NewAnalyzer(apiKey string) *Analyzer {
	client := openai.NewClient(apiKey)
	return &Analyzer{
		client: client,
	}
}

// AnalyzeCode performs analysis on the provided code and returns findings
func (a *Analyzer) AnalyzeCode(ctx context.Context, code, filename string) ([]models.Finding, error) {
	// Create the prompt for OpenAI
	prompt := fmt.Sprintf(`Analyze the following code for security vulnerabilities and code quality issues.
	Return the findings in JSON format with the following structure:
	[
		{
			"type": "vulnerability|code_smell|best_practice",
			"message": "description of the issue",
			"line": line_number,
			"severity": "high|medium|low",
			"file": "%s",
			"code": "relevant code snippet"
		}
	]
	
	Only return issues that are actually present in the code. Do not make up issues.
	
	Code to analyze:
	%s`, filename, code)

	// Call OpenAI API
	resp, err := a.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a code security auditor. Your job is to identify security vulnerabilities, code smells, and best practice violations in code.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error calling OpenAI API: %w", err)
	}

	// Parse the response
	var findings []models.Finding
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &findings); err != nil {
		log.Printf("Error parsing OpenAI response: %v", err)
		log.Printf("Response content: %s", resp.Choices[0].Message.Content)
		return nil, fmt.Errorf("error parsing OpenAI response: %w", err)
	}

	return findings, nil
}

// GetMockFindings returns mock findings for testing purposes
func GetMockFindings(filename string) []models.Finding {
	return []models.Finding{
		{
			Type:     "vulnerability",
			Message:  "SQL injection vulnerability detected",
			Line:     15,
			Severity: "high",
			File:     filename,
			Code:     "db.Query(\"SELECT * FROM users WHERE id = \" + userId)",
		},
		{
			Type:     "code_smell",
			Message:  "Function is too long and complex",
			Line:     30,
			Severity: "medium",
			File:     filename,
			Code:     "func complexFunction() { ... }",
		},
		{
			Type:     "best_practice",
			Message:  "Missing error handling",
			Line:     22,
			Severity: "low",
			File:     filename,
			Code:     "result, _ := someOperation()",
		},
	}
}
