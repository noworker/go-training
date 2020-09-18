package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"go_training/domain/model"
)

func MakeHashedString(s string) model.HashedString {
	r := sha256.Sum256([]byte(s))
	return model.HashedString(hex.EncodeToString(r[:]))
}

