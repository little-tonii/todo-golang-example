package common

type ApplicationError struct {
	StatusCode int
	Error      error
}

func NewApplicationError(statusCode int, error error) *ApplicationError {
	return &ApplicationError{
		StatusCode: statusCode,
		Error:      error,
	}
}
