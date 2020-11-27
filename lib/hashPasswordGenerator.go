package lib

import (
	"go_training/lib/errors"
	"go_training/web/api_error"
	"golang.org/x/crypto/bcrypt"
)

type HashedByteString []byte

const (
	PasswordIsTooShort errors.ErrorMessage = "password is too short"
	PasswordItTooLong  errors.ErrorMessage = "password is too long"
)

const minPasswordLen = 8
const maxPasswordLen = 255
const bcryptCost = 11

func MakeHashedStringFromPassword(s string) (HashedByteString, error) {
	println("MakeHashedStringFromPassword: ", s)
	if len(s) < minPasswordLen {
		return []byte{}, api_error.InvalidRequestError(errors.CustomError{Message: PasswordIsTooShort})
	}
	if len(s) > maxPasswordLen {
		return []byte{}, api_error.InvalidRequestError(errors.CustomError{Message: PasswordItTooLong})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcryptCost)
	if err != nil {
		return []byte{}, api_error.InternalError(err)
	}
	println("MakeHashedStringFromPassword: ", string(hashedPassword))
	return hashedPassword, nil
}
