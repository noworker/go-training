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
	newUser, err := model.NewUser(userId, address)
	if err != nil {
		return err
	}

	password, err := lib.MakeHashedStringFromPassword(rawPassword)
	if err != nil {
		return err
	}

	newUserPassword := model.NewUserPassword(newUser.UserId, password)

	if err := service.UserRepository.CreateUnactivatedNewUser(newUser, newUserPassword); err != nil {
		return err
	}
	return nil
}
