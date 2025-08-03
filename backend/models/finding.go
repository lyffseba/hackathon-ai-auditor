package models

// Finding represents a single code analysis finding
type Finding struct {
	Type     string `json:"type"`
	Message  string `json:"message"`
	Line     int    `json:"line"`
	Severity string `json:"severity"`
	File     string `json:"file"`
	Code     string `json:"code,omitempty"`
}
