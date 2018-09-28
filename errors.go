package main

type NotFoundError struct {
	error
	msg string
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		msg: msg,
	}
}

func (err NotFoundError) Error() string {
	return err.msg
}

type AlreadyExistsError struct {
	error
	msg string
}

func NewAlreadyExistsError(msg string) *AlreadyExistsError {
	return &AlreadyExistsError{
		msg: msg,
	}
}

func (err AlreadyExistsError) Error() string {
	return err.msg
}
