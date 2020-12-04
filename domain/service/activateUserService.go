package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
)

type ActivateUserService struct {
	TokenChecker   infrainterface.ITokenChecker
	UserRepository infrainterface.IUserRepository
}

func NewActivateUserService(checker infrainterface.ITokenChecker, repository infrainterface.IUserRepository) ActivateUserService {
	return ActivateUserService{TokenChecker: checker, UserRepository: repository}
}

func (service ActivateUserService) ActivateUser(token model.Token) error {
	userId, err := service.TokenChecker.CheckActivateUserToken(token)
	if err != nil {
		return err
	}
	if err := service.UserRepository.Activate(userId); err != nil {
		return err
	}
	return nil
}
