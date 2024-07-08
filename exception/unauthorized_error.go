package exception

type UnAuthorizedError struct {
	Message string `json:"message"`
}

func (ce *UnAuthorizedError) Error() string {
	return ce.Message
}

func NewUnAuthorizedError(message string) *UnAuthorizedError {
	return &UnAuthorizedError{
		Message: message,
	}
}