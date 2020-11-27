package jw_token

import (
	"github.com/dgrijalva/jwt-go"
	mdate "github.com/matsuri-tech/date-go"
	"go_training/lib/errors"
	"io/ioutil"
	"strconv"
)

const (
	NoFileError          errors.ErrorMessage = "no file"
	ParsePrivateKeyError errors.ErrorMessage = "parse private privateKey error"
	EncodeTokenError     errors.ErrorMessage = "encode token error"
)

type TokenGenerator struct {
	privateKeyPath []byte
}

func NewTokenGenerator(path string) (TokenGenerator, error) {
	signBytes, err := ioutil.ReadFile(path + "private.pem")
	if err != nil {
		return TokenGenerator{}, errors.CustomError{Message: NoFileError, Option: err.Error()}
	}
	return TokenGenerator{privateKeyPath: signBytes}, nil
}

func (g TokenGenerator) GenerateActivateUserToken(userId string) (string, error) {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(g.privateKeyPath)
	if err != nil {
		return "", errors.CustomError{Message: ParsePrivateKeyError, Option: err.Error()}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     strconv.Itoa(int(mdate.GetToday().PlusNDay(1).Unix())),
	})

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", errors.CustomError{Message: EncodeTokenError, Option: err.Error()}
	}

	return tokenString, nil
}
