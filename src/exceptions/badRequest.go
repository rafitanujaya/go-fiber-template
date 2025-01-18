package exceptions

type BadRequestError struct {
	Message    string
	StatusCode int16
}

func NewBadRequestError(message string, statusCode int16) *BadRequestError {
	return &BadRequestError{Message: message, StatusCode: statusCode}
}

func (e *BadRequestError) Error() string {
	return e.Message
}
