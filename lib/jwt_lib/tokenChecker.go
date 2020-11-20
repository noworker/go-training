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
	InvalidTokenFormat      errors.ErrorMessage = "invalid token format"
)

func Checker(jwtStr string, conf config.Config) (string, error) {
	verifyBytes, err := ioutil.ReadFile(conf.App.KeyPath + "public.pem")
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.CustomError{Message: UnexpectedSigningMethod}
		} else {
			return signKey, nil
		}
	})
	if err != nil {
		return "", err
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", errors.CustomError{Message: InvalidTokenFormat}
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.CustomError{Message: InvalidTokenFormat}
		}
		return userId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.CustomError{Message: NotEvenAToken}
		} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			return "", errors.CustomError{Message: Expired}
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
			return "", errors.CustomError{Message: NotValidYet}
		} else {
			return "", errors.CustomError{Message: CanNotHandle}
		}
	} else {
		return "", errors.CustomError{Message: CanNotHandle}
	}
}
