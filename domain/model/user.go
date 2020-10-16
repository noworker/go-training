package model

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/badoux/checkmail"
	"go_training/lib/errors"
)

const (
	InvalidEmailAddressFormat errors.ErrorMessage = "invalid email address format"
	UserIdNotExists           errors.ErrorMessage = "user_id not exists"
	EmailAddressNotExists     errors.ErrorMessage = "email_address not exists"
)

const (
	PasswordNotExists  errors.ErrorMessage = "password not exists"
	PasswordIsTooShort errors.ErrorMessage = "password is too short"
)

type UserId string
type EmailAddress string
type HashString string

type User struct {
	UserId
	EmailAddress
	Password  HashString
	Activated bool
}

func NewUser(userIdString, emailAddressString, passwordString string) (User, error) {
	userId, err := newUserId(userIdString)
	if err != nil {
		return User{}, err
	}

	emailAddress, err := newEmailAddress(emailAddressString)
	if err != nil {
		return User{}, err
	}
	password, err := newPassword(passwordString)
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

func newPassword(password string) (HashString, error) {
	if len(password) < 8 {
		return "", errors.CustomError{Message: PasswordIsTooShort}
	}
	return MakeHashedStringFromPassword(password), nil
}

func MakeHashedStringFromPassword(s string) HashString {
	r := sha256.Sum256([]byte(s))
	return HashString(hex.EncodeToString(r[:]))
}
