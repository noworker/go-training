package repository

import (
	"go_training/domain/model"
	"go_training/lib/errors"
)

type UserRepositoryMock struct {
	ExistingUserId model.UserId
	User           model.User
	UserPassword   model.UserPassword
}

func NewUserRepositoryMock(existingUserId string) *UserRepositoryMock {
	return &UserRepositoryMock{
		ExistingUserId: model.UserId(existingUserId),
	}
}

func (repository *UserRepositoryMock) Activate(userId model.UserId) error {
	if repository.User.UserId != userId {
		panic("userId or password does not match")
	}
	repository.User.Activated = true
	return nil
}

func (repository *UserRepositoryMock) CreateUser(user model.User, userPassword model.UserPassword) error {
	if user.UserId == repository.ExistingUserId {
		return errors.CustomError{Message: CanNotCreateExistingUserId}
	}
	repository.User = user
	repository.UserPassword = userPassword
	return nil
}
