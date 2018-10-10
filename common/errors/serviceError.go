package errors

type ServiceError struct {
	message string
}

func NewServiceError() *ServiceError {
	return &ServiceError{
		message: "internal service error",
	}
}

func (err *ServiceError) Error() string {
	return err.message
}
