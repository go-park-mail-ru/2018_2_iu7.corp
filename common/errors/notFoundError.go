package errors

type NotFoundError struct {
	ServiceError
}

func NewNotFoundError(message string) *NotFoundError {
	err := &NotFoundError{}
	err.message = message
	return err
}
