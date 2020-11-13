package lib

import (
	"github.com/dgrijalva/jwt-go"
	"go_training/lib/errors"
	"io/ioutil"
)

const (
	JWTUnexpectedSigningMethod errors.ErrorMessage = "jwt unexpected signing method"
	JWTNotEvenAToken           errors.ErrorMessage = "not even a token"
	JWTExpired                 errors.ErrorMessage = "token is expired"
	JWTNotValidYet             errors.ErrorMessage = "token is not valid yet"
	JWTCanNotHandle            errors.ErrorMessage = "can not handle this token"
)

func JWTChecker(jwtStr JWTStr) (bool, error) {
	verifyBytes, err := ioutil.ReadFile("./demo.rsa.pub.pkcs8")
	if err != nil {
		return false, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return false, err
	}

	token, err := jwt.Parse(string(jwtStr), func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, errors.CustomError{Message: JWTUnexpectedSigningMethod}
		} else {
			return verifyKey, nil
		}
	})
	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.CustomError{Message: JWTNotEvenAToken}
		} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			return false, errors.CustomError{Message: JWTExpired}
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.CustomError{Message: JWTNotValidYet}
		} else {
			return false, errors.CustomError{Message: JWTCanNotHandle}
		}
	} else {
		return false, errors.CustomError{Message: JWTCanNotHandle}
	}
}
