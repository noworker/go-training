package repository

import (
	"go_training/domain/model"
	"go_training/lib/errors"
)

type UserRepositoryMock struct {
	userValue
}

type userValue struct {
	ExistingUserId model.UserId
	UserId model.UserId
	EmailAddress model.EmailAddress
	Password model.HashString
	Activated bool
}

func NewUserRepositoryMock(existingUserId string) *UserRepositoryMock {
	return  &UserRepositoryMock{
		userValue{ExistingUserId: model.UserId(existingUserId)},
	}
}

func (repository *UserRepositoryMock) Activate(userId model.UserId, password model.HashString) error {
	if repository.UserId != userId || repository.Password != password {
		panic("userId or password does not match")
	}
	repository.userValue.Activated = true
	return nil
}

func (repository *UserRepositoryMock) CreateUnactivatedNewUser(userId model.UserId, emailAddress model.EmailAddress, password model.HashString) error {
	if userId == repository.ExistingUserId {
		return errors.CustomError{Message: "can not create existing user id", ErrorType: errors.CanNotCreateExistingUserId}
	}
	repository.userValue = userValue{UserId: userId, EmailAddress: emailAddress, Password: password, Activated: false}
	return nil
}
