package main

type ServerError struct {
	error
	msg string
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

func NewNotFoundError(msg string) *NotFoundError {
	err := &NotFoundError{}
	err.msg = msg
	return err
}

func (err NotFoundError) Error() string {
	return err.msg
}

func NewAlreadyExistsError(msg string) *AlreadyExistsError {
	err := &AlreadyExistsError{}
	err.msg = msg
	return err
}

func (err AlreadyExistsError) Error() string {
	return err.msg
}

func NewInvalidFormatError(msg string) *InvalidFormatError {
	err := &InvalidFormatError{}
	err.msg = msg
	return err
}

func (err AlreadyExistsError) Error() string {
	return err.msg
}
