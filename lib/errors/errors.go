package errors

import "fmt"

type CustomError struct {
	Message ErrorMessage
	Option  string
}

type ErrorMessage string

func (err CustomError) Error() string {
	if err.Option == "" {
		return fmt.Sprintf("%s", err.Message)
	}
	return fmt.Sprintf("%s, %s", err.Message, err.Option)
}
