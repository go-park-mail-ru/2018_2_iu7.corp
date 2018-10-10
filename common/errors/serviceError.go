package errors

type ServiceError struct {
	message string
}

func NewServiceError() *ServiceError {
	return &ServiceError{
		message: "internal services error",
	}
}

func (err *ServiceError) Error() string {
	return err.message
}
