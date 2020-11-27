package jw_token

import (
	"github.com/dgrijalva/jwt-go"
	"go_training/config"
	"go_training/lib/errors"
	"io/ioutil"
	"strconv"
	"time"
)

const (
	UnexpectedSigningMethod errors.ErrorMessage = "unexpected signing method"
	NotEvenAToken           errors.ErrorMessage = "not even a token"
	Expired                 errors.ErrorMessage = "token is expired"
	NotValidYet             errors.ErrorMessage = "token is not valid yet"
	CanNotHandle            errors.ErrorMessage = "can not handle this token"
	InvalidTokenFormat      errors.ErrorMessage = "invalid token format"
	ParseTokenError         errors.ErrorMessage = "parse token error"
	ParsePublicKeyError     errors.ErrorMessage = "parse public key error"
	TokenIsExpired          errors.ErrorMessage = "token is expired"
)

func isExpired(exp int) bool {
	now := time.Now().Unix()
	if exp < int(now) {
		return true
	}
	return false
}

func TokenChecker(jwtStr string, conf config.Config) (string, error) {
	verifyBytes, err := ioutil.ReadFile(conf.App.KeyPath + "public.pem")
	if err != nil {
		return "", errors.CustomError{Message: NoFileError, Option: err.Error()}
	}

	signKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return "", errors.CustomError{Message: ParsePublicKeyError, Option: err.Error()}
	}

	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.CustomError{Message: UnexpectedSigningMethod, Option: err.Error()}
		} else {
			return signKey, nil
		}
	})
	if err != nil {
		return "", errors.CustomError{Message: ParseTokenError, Option: err.Error()}
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", errors.CustomError{Message: InvalidTokenFormat}
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get user_id from token"}
		}
		expiredAt, ok := claims["exp"].(string)
		if !ok {
			return "", errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get exp from token"}
		}
		exp, err := strconv.Atoi(expiredAt)
		if err != nil {
			return "", errors.CustomError{Message: InvalidTokenFormat, Option: err.Error()}
		}
		if isExpired(exp) {
			return "", errors.CustomError{Message: TokenIsExpired}
		}
		return userId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.CustomError{Message: NotEvenAToken, Option: ve.Error()}
		} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			return "", errors.CustomError{Message: Expired, Option: ve.Error()}
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
			return "", errors.CustomError{Message: NotValidYet, Option: ve.Error()}
		} else {
			return "", errors.CustomError{Message: CanNotHandle, Option: ve.Error()}
		}
	} else {
		return "", errors.CustomError{Message: CanNotHandle, Option: ve.Error()}
	}
}
