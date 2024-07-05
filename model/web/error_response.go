package web

type ErrorResponse struct {
	Errors []DetailError `json:"errors"`
}

type DetailError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}
