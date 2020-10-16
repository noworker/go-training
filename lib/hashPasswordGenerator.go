package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"go_training/lib/errors"
)

type HashString string

const (
	PasswordIsTooShort errors.ErrorMessage = "password is too short"
)

func MakeHashedStringFromPassword(s string) (HashString, error) {
	if len(s) < 8 {
		return "", errors.CustomError{Message: PasswordIsTooShort}
	}
	r := sha256.Sum256([]byte(s))
	return HashString(hex.EncodeToString(r[:])), nil
}
