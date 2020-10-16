package errors

type CustomError struct {
	Message ErrorMessage
}

type ErrorMessage string

func (err CustomError) Error() string {
	return string(err.Message)
}
