package jw_token

import "go_training/domain/model"

type TokenGeneratorMock struct {
}

func NewTokenGeneratorMock(path string) (TokenGeneratorMock, error) {
	return TokenGeneratorMock{}, nil
}

func (g TokenGeneratorMock) GenerateActivateUserToken(userId model.UserId) (model.Token, error) {
	return "token", nil
}
