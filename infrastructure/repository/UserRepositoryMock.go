package repository

import (
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib"
	"go_training/lib/errors"
	"go_training/web/api_error"
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

func (repository *UserRepositoryMock) CreateNewUser(user model.User, userPassword model.UserPassword) error {
	if user.UserId == repository.ExistingUserId {
		return api_error.InvalidRequestError(errors.CustomError{Message: CanNotCreateExistingUserId})
	}
	repository.User = user
	repository.UserPassword = userPassword
	return nil
}

func (repository UserRepositoryMock) UserExists(userId model.UserId, password lib.HashedByteString) bool {
	return true
}

func (repository UserRepositoryMock) GetUserByIdAndPassword(userId model.UserId, password lib.HashedByteString) (table.User, error) {
	return table.User{}, nil
}
