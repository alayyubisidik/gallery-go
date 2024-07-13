package exception

type BadRequestError struct {
	Message string `json:"message"`
}

func (ce *BadRequestError) Error() string {
	return ce.Message
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
	}
}