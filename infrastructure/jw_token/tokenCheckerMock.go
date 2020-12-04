package jw_token

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib/errors"
	"go_training/web/api_error"
)

type TokenCheckerMock struct {
	model.UserId
	Token model.Token
}

func NewTokenCheckerMock(id model.UserId, token string) (infrainterface.ITokenChecker, error) {
	return TokenCheckerMock{UserId: id, Token: model.Token(token)}, nil
}

func (c TokenCheckerMock) CheckActivateUserToken(jwtStr model.Token) (model.UserId, error) {
	if c.Token != jwtStr {
		return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get user_id from token"})
	}
	return c.UserId, nil
}

func (c TokenCheckerMock) CheckLoginUserToken(jwtStr model.Token) (model.UserId, error) {
	if c.Token != jwtStr {
		return "", api_error.InvalidRequestError(errors.CustomError{Message: InvalidTokenFormat, Option: "failed to get user_id from token"})
	}
	return c.UserId, nil
}
