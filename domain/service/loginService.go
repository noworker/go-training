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

func (service LoginService) Login(userId, password string) error {
	user, err := service.UserRepository.GetUserByIdAndPassword(model.UserId(userId), password)
	if err != nil {
		return err
	}

	token, err := service.TokenGenerator.GenerateTwoStepVerificationToken(model.UserId(userId))
	if err != nil {
		return err
	}

	go service.EmailSender.SendTwoStepVerificationEmail(user.EmailAddress, token)
	return nil
}
