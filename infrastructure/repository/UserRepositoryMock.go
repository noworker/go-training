package repository

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
)

type userRepositoryMock struct {
	userValue
}

type userValue struct {
	UserId model.UserId
	EmailAddress model.EmailAddress
	Password model.HashString
	Activated bool
}

func NewUserRepositoryMock() infrainterface.IUserRepository {
	return  userRepositoryMock{
	}
}

func (repository userRepositoryMock) Activate(userId model.UserId, password model.HashString) error {
	if repository.UserId != userId || repository.Password != password {
		panic("userId or password does not match")
	}
	repository.userValue.Activated = true
	return nil
}

func (repository userRepositoryMock) CreateUnactivatedNewUser(userId model.UserId, emailAddress model.EmailAddress, password model.HashString) error {
	repository.userValue = userValue{UserId: userId, EmailAddress: emailAddress, Password: password, Activated: false}
	return nil
}
