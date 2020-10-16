package validator

import (
	"github.com/badoux/checkmail"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

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

func ValidateLength(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return len(str) >= 8
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.RegisterValidation("email_address", ValidateEmailAddress); err != nil {
		return err
	}
	if err := cv.validator.RegisterValidation("password", ValidateLength); err != nil {
		return err
	}
	return cv.validator.Struct(i)
}