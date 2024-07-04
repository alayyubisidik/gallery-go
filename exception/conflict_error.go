package exception

type ConflictError struct{
	Error string
}

func NewConflictError(err string) ConflictError {
	return ConflictError{
		Error: err,
	}
}