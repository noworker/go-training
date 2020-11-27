package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib"
	"go_training/lib/errors"
	"go_training/web/api_error"
)

const (
	NoUserError                 errors.ErrorMessage = "no_user_error"
	UserIsAlreadyActivatedError errors.ErrorMessage = "user_is_already_activated_error"
)

type ResendActivationEmailService struct {
	UserRepository infrainterface.IUserRepository
	TokenGenerator infrainterface.ITokenGenerator
	EmailSender    infrainterface.IEmail
}

func NewResendActivationEmailService(userRepository infrainterface.IUserRepository,
	tokenGenerator infrainterface.ITokenGenerator,
	emailSender infrainterface.IEmail) ResendActivationEmailService {
	return ResendActivationEmailService{UserRepository: userRepository, TokenGenerator: tokenGenerator, EmailSender: emailSender}
}

func (service ResendActivationEmailService) ResendActivationEmail(userId, password, address string) error {
	hashedPassword, err := lib.MakeHashedStringFromPassword(password)
	if err != nil {
		return api_error.InvalidRequestError(err)
	}

	err = service.checkIfUserIsActivated(model.UserId(userId), hashedPassword)
	if err != nil {
		return err
	}

	token, err := service.TokenGenerator.GenerateActivateUserToken(userId)
	if err != nil {
		return err
	}

	go service.EmailSender.SendEmail(address, token)
	return nil
}

func (service ResendActivationEmailService) checkIfUserIsActivated(userId model.UserId, hashedPassword lib.HashedByteString) error {
	exists := service.UserRepository.UserExists(userId, hashedPassword)
	if !exists {
		return api_error.InvalidRequestError(errors.CustomError{Message: NoUserError})
	}

	user, err := service.UserRepository.GetUserByIdAndPassword(userId, hashedPassword)
	if err != nil {
		return api_error.InvalidRequestError(err)
	}

	if user.Activated {
		return api_error.InvalidRequestError(errors.CustomError{Message: UserIsAlreadyActivatedError})
	}

	return nil
}
