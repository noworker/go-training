package model

import (
	"github.com/badoux/checkmail"
	"go_training/lib"
	"go_training/lib/errors"
)

const (
	InvalidEmailAddressFormat errors.ErrorMessage = "invalid email address format"
	UserIdNotExists           errors.ErrorMessage = "user_id not exists"
	EmailAddressNotExists     errors.ErrorMessage = "email_address not exists"
)

type UserId string
type EmailAddress string

type User struct {
	UserId
	EmailAddress
	Password  lib.HashString
	Activated bool
}

func NewUser(userIdString, emailAddressString string, password lib.HashString) (User, error) {
	userId, err := newUserId(userIdString)
	if err != nil {
		return User{}, err
	}

	emailAddress, err := newEmailAddress(emailAddressString)
	if err != nil {
		return User{}, err
	}

	return User{UserId: userId, EmailAddress: emailAddress, Password: password}, nil
}

func newUserId(userId string) (UserId, error) {
	if userId == "" {
		return UserId(0), errors.CustomError{Message: UserIdNotExists}
	}
	return UserId(userId), nil
}

func newEmailAddress(emailAddress string) (EmailAddress, error) {
	if err := checkmail.ValidateFormat(emailAddress); err != nil {
		println("emailError, ", emailAddress, err.Error())
		return "", err
	}
	return EmailAddress(emailAddress), nil
}
