package errors

type ServerError struct {
	error
	msg string
}

type AlreadyAuthorizedError struct {
	ServerError
}

type NotFoundError struct {
	ServerError
}

type AlreadyExistsError struct {
	ServerError
}

type InvalidFormatError struct {
	ServerError
}

func NewAlreadyAuthorizedError(msg string) *AlreadyAuthorizedError {
	err := &AlreadyAuthorizedError{}
	err.msg = msg
	return err
}

func NewNotFoundError(msg string) *NotFoundError {
	err := &NotFoundError{}
	err.msg = msg
	return err
}

func NewAlreadyExistsError(msg string) *AlreadyExistsError {
	err := &AlreadyExistsError{}
	err.msg = msg
	return err
}

func NewInvalidFormatError(msg string) *InvalidFormatError {
	err := &InvalidFormatError{}
	err.msg = msg
	return err
}

func (err ServerError) Error() string {
	return err.msg
}
