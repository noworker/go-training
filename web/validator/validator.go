package validator

import (
	"github.com/badoux/checkmail"
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func ValidateEmailAddress(fl validator.FieldLevel) bool {
	emailAddress := fl.Field().String()
	if err := checkmail.ValidateFormat(emailAddress); err != nil {
		return false
	}
	return true
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.RegisterValidation("email_address", ValidateEmailAddress); err != nil {
		return err
	}
	return cv.validator.Struct(i)
}