package errors

type ConstraintViolationError struct {
	ServiceError
}

func NewConstraintViolationError(message string) *ConstraintViolationError {
	err := &ConstraintViolationError{}
	err.message = message
	return err
}
