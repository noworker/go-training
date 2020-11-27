package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib"
)

type CreateUserService struct {
	UserRepository infrainterface.IUserRepository
	TokenGenerator infrainterface.ITokenGenerator
	EmailSender    infrainterface.IEmail
}

func NewCreateUserService(userRepository infrainterface.IUserRepository, tokenGenerator infrainterface.ITokenGenerator, emailSender infrainterface.IEmail) CreateUserService {
	return CreateUserService{
		UserRepository: userRepository,
		TokenGenerator: tokenGenerator,
		EmailSender:    emailSender,
	}
}

func (service CreateUserService) CreateUser(userId string, address string, rawPassword string) error {
	newUser, err := model.NewUser(userId, address)
	if err != nil {
		return err
	}

	password, err := lib.MakeHashedStringFromPassword(rawPassword)
	if err != nil {
		return err
	}

	newUserPassword := model.NewUserPassword(newUser.UserId, password)

	if err := service.UserRepository.CreateUser(newUser, newUserPassword); err != nil {
		return err
	}

	token, err := service.TokenGenerator.GenerateActivateUserToken(userId)
	if err != nil {
		return err
	}

	err = service.EmailSender.SendEmail(address, token)
	if err != nil {
		return err
	}

	return nil
}
