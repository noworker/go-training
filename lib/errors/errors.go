package errors

type CustomError struct {
	Message string
	ErrorType ErrorType
}

type ErrorType string

func (err CustomError) Error() string {
	return err.Message
}

const (
	CanNotCreateExistingUserId ErrorType = "can_not_create_existing_user_id"
)