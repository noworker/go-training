package service

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib"
)

type CreateUserService struct {
	UserRepository infrainterface.IUserRepository
}

func NewCreateUserService(userRepository infrainterface.IUserRepository) CreateUserService {
	return CreateUserService{
		UserRepository: userRepository,
	}
}

func (service CreateUserService) CreateUser(userId string, address string, rawPassword string) error {
	password, err := lib.MakeHashedStringFromPassword(rawPassword)
	if err != nil {
		return err
	}

	newUser, err := model.NewUser(userId, address, password)
	if err != nil {
		return err
	}

	token := lib.MakeUniqueToken()

	if err := service.UserRepository.CreateUnactivatedNewUser(newUser, token); err != nil {
		return err
	}
	return nil
}
