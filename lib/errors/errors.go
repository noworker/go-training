package errors

type CustomError struct {
	Message ErrorMessage
}

type ErrorMessage string

func (err CustomError) Error() string {
	return string(err.Message)
}
const (
	CanNotCreateExistingUserId ErrorMessage = "can_not_create_existing_user_id"
)