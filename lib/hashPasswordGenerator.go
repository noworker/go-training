package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"go_training/lib/errors"
)

type HashString string

const (
	PasswordIsTooShort errors.ErrorMessage = "password is too short"
	PasswordItTooLong  errors.ErrorMessage = "password is too long"
)

const minPasswordLen = 8
const maxPasswordLen = 255

func MakeHashedStringFromPassword(s string) (HashString, error) {
	if len(s) < minPasswordLen {
		return "", errors.CustomError{Message: PasswordIsTooShort}
	}
	if len(s) > maxPasswordLen {
		return "", errors.CustomError{Message: PasswordItTooLong}
	}
	r := sha256.Sum256([]byte(s))
	return HashString(hex.EncodeToString(r[:])), nil
}
