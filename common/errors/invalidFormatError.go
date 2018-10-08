package errors

type InvalidFormatError struct {
	ServiceError
}

func NewInvalidFormatError(message string) *InvalidFormatError {
	err := &InvalidFormatError{}
	err.message = message
	return err
}
