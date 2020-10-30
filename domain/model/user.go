package model

import (
	"github.com/badoux/checkmail"
	"go_training/lib"
	"go_training/lib/errors"
)

const (
	UserIdNotExists           errors.ErrorMessage = "user_id not exists"
	UserIdIsTooLong           errors.ErrorMessage = "user is is too long"
	InvalidEmailAddressFormat errors.ErrorMessage = "invalid email address format"
)

const maxUserIdLen = 16

type UserId string
type EmailAddress string

type User struct {
	UserId
	EmailAddress
	Password  lib.HashedByteString
	Activated bool
}

func NewUser(userIdString, emailAddressString string, password lib.HashedByteString) (User, error) {
	userId, err := newUserId(userIdString)
	if err != nil {
		return User{}, err
	}

	emailAddress, err := newEmailAddress(emailAddressString)
	if err != nil {
		return User{}, errors.CustomError{Message: InvalidEmailAddressFormat}
	}

	return User{UserId: userId, EmailAddress: emailAddress, Password: password}, nil
}

func newUserId(userId string) (UserId, error) {
	if userId == "" {
		return "", errors.CustomError{Message: UserIdNotExists}
	}
	if len(userId) > maxUserIdLen {
		return "", errors.CustomError{Message: UserIdIsTooLong}
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
