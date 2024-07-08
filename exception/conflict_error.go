package exception

type ConflictError struct {
	Message string `json:"message"`
}

func (ce *ConflictError) Error() string {
	return ce.Message
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}