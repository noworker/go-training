package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
)

type LoginService struct {
	UserRepository infrainterface.IUserRepository
	TokenGenerator infrainterface.ITokenGenerator
	EmailSender    infrainterface.IEmail
}

func (service LoginService) Login(userId, password string) (model.Token, error) {
	if _, err := service.UserRepository.GetUserByIdAndPassword(model.UserId(userId), password); err != nil {
		return "", err
	}

	return service.TokenGenerator.GenerateLoginUserToken(model.UserId(userId))
}
