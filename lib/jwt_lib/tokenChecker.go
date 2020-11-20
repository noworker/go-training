package jwt_lib

import (
	"github.com/dgrijalva/jwt-go"
	"go_training/config"
	"go_training/lib/errors"
	"io/ioutil"
)

const (
	UnexpectedSigningMethod errors.ErrorMessage = "unexpected signing method"
	NotEvenAToken           errors.ErrorMessage = "not even a token"
	Expired                 errors.ErrorMessage = "token is expired"
	NotValidYet             errors.ErrorMessage = "token is not valid yet"
	CanNotHandle            errors.ErrorMessage = "can not handle this token"
)

func Checker(jwtStr string, conf config.Config) (bool, error) {
	verifyBytes, err := ioutil.ReadFile(conf.App.KeyPath + "public.pem")
	if err != nil {
		return false, err
	}

	signKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return false, err
	}

	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.CustomError{Message: UnexpectedSigningMethod}
		} else {
			return signKey, nil
		}
	})
	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.CustomError{Message: NotEvenAToken}
		} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			return false, errors.CustomError{Message: Expired}
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.CustomError{Message: NotValidYet}
		} else {
			return false, errors.CustomError{Message: CanNotHandle}
		}
	} else {
		return false, errors.CustomError{Message: CanNotHandle}
	}
}
