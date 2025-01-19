package exceptions

type ConflictError struct {
	Message    string
	StatusCode int
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{Message: message, StatusCode: 409}
}

func (e *ConflictError) Error() string {
	return e.Message
}
