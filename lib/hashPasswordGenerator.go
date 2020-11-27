package lib

import (
	"go_training/lib/errors"
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
	if len(s) < minPasswordLen {
		return []byte{}, errors.CustomError{Message: PasswordIsTooShort}
	}
	if len(s) > maxPasswordLen {
		return []byte{}, errors.CustomError{Message: PasswordItTooLong}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcryptCost)
	if err != nil {
		return []byte{}, err
	}
	return hashedPassword, nil
}
