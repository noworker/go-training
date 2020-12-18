package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
)

type TwoStepVerificationService struct {
	TokenChecker   infrainterface.ITokenChecker
	TokenGenerator infrainterface.ITokenGenerator
}

func (service TwoStepVerificationService) Verify(token model.Token) (model.Token, error) {
	userID, err := service.TokenChecker.CheckTwoStepVerificationToken(token)
	if err != nil {
		return "", err
	}

	loginToken, err := service.TokenGenerator.GenerateLoginUserToken(userID)
	if err != nil {
		return "", err
	}

	return loginToken, nil
}
