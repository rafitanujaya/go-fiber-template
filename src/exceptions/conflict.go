package exceptions

type ConflictError struct {
	Message    string
	StatusCode int16
}

func NewConflictError(message string, statusCode int16) *ConflictError {
	return &ConflictError{Message: message, StatusCode: statusCode}
}

func (e *ConflictError) Error() string {
	return e.Message
}
