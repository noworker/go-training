package jw_token

import (
	"github.com/dgrijalva/jwt-go"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib/errors"
	"go_training/web/api_error"
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
	ParsePublicKeyError     errors.ErrorMessage = "parse public privateKey error"
	TokenIsExpired          errors.ErrorMessage = "token is expired"
)

func isExpired(exp int) bool {
	now := time.Now().Unix()
	if exp < int(now) {
		return true
	}
	return false
}

type TokenChecker struct {
	publicKey []byte
}

func NewTokenChecker(path string) (infrainterface.ITokenChecker, error) {
	pubKey, err := ioutil.ReadFile(path + "public.pem")
	if err != nil {
		return TokenChecker{}, errors.CustomError{Message: NoFileError, Option: err.Error()}
	}
	return TokenChecker{publicKey: pubKey}, nil
}

func (c TokenChecker) checkToken(jwtStr model.Token, tokenType TokenType) (model.UserId, error) {
	signKey, err := jwt.ParseRSAPublicKeyFromPEM(c.publicKey)
	if err != nil {
		return "", api_error.InternalError(errors.CustomError{Message: ParsePublicKeyError, Option: err.Error()})
	}

	token, err := jwt.Parse(string(jwtStr), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, api_error.InvalidRequestError(errors.CustomError{Message: UnexpectedSigningMethod, Option: err.Error()})
		} else {
			return signKey, nil
		}
	})
	if err != nil {
		return "", api_error.InvalidRequestError(errors.CustomError{Message: ParseTokenError, Option: err.Error()})
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat})
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get user_id from token"})
		}
		expiredAt, ok := claims["exp"].(string)
		if !ok {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get exp from token"})
		}
		tokenTypeFromJWT, ok := claims["type"].(string)
		if !ok {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get token type from token"})
		}
		if string(tokenType) != tokenTypeFromJWT {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get appropriate token type from token"})
		}
		exp, err := strconv.Atoi(expiredAt)
		if err != nil {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: err.Error()})
		}
		if isExpired(exp) {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: TokenIsExpired})
		}
		return model.UserId(userId), nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.CustomError{Message: NotEvenAToken, Option: ve.Error()}
		} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			return "", errors.CustomError{Message: Expired, Option: ve.Error()}
		} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: NotValidYet, Option: ve.Error()})
		} else {
			return "", api_error.InvalidRequestError(errors.CustomError{Message: CanNotHandle, Option: ve.Error()})
		}
	} else {
		return "", api_error.InvalidRequestError(errors.CustomError{Message: CanNotHandle, Option: ve.Error()})
	}
}

func (c TokenChecker) CheckActivateUserToken(jwtStr model.Token) (model.UserId, error) {
	return c.checkToken(jwtStr, Activation)
}

func (c TokenChecker) CheckLoginUserToken(jwtStr model.Token) (model.UserId, error) {
	return c.checkToken(jwtStr, Login)
}
