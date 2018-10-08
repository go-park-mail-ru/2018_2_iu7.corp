package errors

type DuplicateError struct {
	ServiceError
}

func NewDuplicateError(message string) *DuplicateError {
	err := &DuplicateError{}
	err.message = message
	return err
}
