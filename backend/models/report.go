package models

import (
	"time"
)

// Report represents a code analysis report
type Report struct {
	ID         string    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	Repository string    `json:"repository,omitempty"`
	Findings   []Finding `json:"findings"`
	Summary    Summary   `json:"summary"`
}

// Summary contains aggregated information about the findings in a report
type Summary struct {
	TotalFindings          int `json:"total_findings"`
	SecurityIssues         int `json:"security_issues"`
	PerformanceIssues      int `json:"performance_issues"`
	BestPracticeViolations int `json:"best_practice_violations"`
}
