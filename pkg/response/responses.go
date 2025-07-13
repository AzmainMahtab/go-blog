// Package responses manages API responses in a standardized format.
package responses

// Response is the standardized API response structure
type Response struct {
	Success    bool        `json:"success"`
	Status     string      `json:"status"`            // "success", "fail", or "error"
	StatusCode int         `json:"statusCode"`        // HTTP status code
	Message    string      `json:"message,omitempty"` // Human-readable message
	Data       interface{} `json:"data,omitempty"`    // Primary response data
	Meta       interface{} `json:"meta,omitempty"`    // Pagination, etc.
	Errors     []ErrorItem `json:"errors,omitempty"`  // Error details
}

// ErrorItem represents detailed error information
type ErrorItem struct {
	Code    string `json:"code,omitempty"`    // Machine-readable code
	Field   string `json:"field,omitempty"`   // For field-specific errors
	Message string `json:"message,omitempty"` // Error description
}
