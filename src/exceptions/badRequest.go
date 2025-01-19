package exceptions

type BadRequestError struct {
	Message    string
	StatusCode int
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{Message: message, StatusCode: 400}
}

func (e *BadRequestError) Error() string {
	return e.Message
}
