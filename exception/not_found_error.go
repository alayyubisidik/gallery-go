package exception

type NotFoundError struct {
	Message string `json:"message"`
}

func (ce *NotFoundError) Error() string {
	return ce.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}