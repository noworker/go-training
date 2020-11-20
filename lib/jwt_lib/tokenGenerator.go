package jwt_lib

import (
	"github.com/dgrijalva/jwt-go"
	mdate "github.com/matsuri-tech/date-go"
	"go_training/config"
	"go_training/lib/errors"
	"io/ioutil"
	"strconv"
)

const (
	NoFileError          errors.ErrorMessage = "no file"
	ParsePrivateKeyError errors.ErrorMessage = "parse private key error"
	EncodeTokenError     errors.ErrorMessage = "encode token error"
)

func TokenGenerator(userId string, conf config.Config) (string, error) {
	signBytes, err := ioutil.ReadFile(conf.App.KeyPath + "private.pem")
	if err != nil {
		return "", errors.CustomError{Message: NoFileError, Option: err.Error()}
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
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
