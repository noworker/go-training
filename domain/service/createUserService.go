package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib"
	"go_training/web/api_error"
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
		return api_error.InvalidRequestError(err)
	}

	password, err := lib.MakeHashedStringFromPassword(rawPassword)
	if err != nil {
		return err
	}

	if err := service.UserRepository.CreateNewUser(newUser, rawPassword, password); err != nil {
		return err
	}
	return nil
}

func (service CreateUserService) SendTokenMail(userId, address string) error {
	token, err := service.TokenGenerator.GenerateActivateUserToken(model.UserId(userId))
	if err != nil {
		return api_error.InternalError(err)
	}

	go service.EmailSender.SendActivationEmail(model.EmailAddress(address), token)

	return nil
}
