package repository

import (
	"go_training/domain/model"
	"go_training/lib"
	"go_training/lib/errors"
	"go_training/web/api_error"
)

type UserRepositoryMock struct {
	UserId       model.UserId
	Password     string
	EmailAddress model.EmailAddress
	Activated    bool
	User         model.User
	UserPassword model.UserPassword
}

func NewUserRepositoryMock(userId, password, address string, activated bool) *UserRepositoryMock {
	return &UserRepositoryMock{
		UserId:       model.UserId(userId),
		Password:     password,
		EmailAddress: model.EmailAddress(address),
		Activated:    activated,
	}
}

func (repository UserRepositoryMock) Activate(userId model.UserId) error {
	if repository.Activated {
		return api_error.InvalidRequestError(errors.CustomError{Message: UserIsAlreadyActivatedError})
	}
	if repository.UserId != userId {
		return api_error.InvalidRequestError(errors.CustomError{Message: UserNotFoundError})
	}

	repository.User.Activated = true

	return nil
}

func (repository *UserRepositoryMock) CreateNewUser(user model.User, rawPassword string, hashedPassword lib.HashedByteString) error {
	if user.UserId == repository.UserId {
		return api_error.InvalidRequestError(errors.CustomError{Message: CanNotCreateExistingUserId})
	}
	repository.User = user
	repository.UserPassword = model.UserPassword{Password: hashedPassword, UserId: user.UserId}
	return nil
}

func (repository UserRepositoryMock) UserExists(userId model.UserId, password lib.HashedByteString) bool {
	return true
}

func (repository UserRepositoryMock) GetUserById(userId model.UserId) (model.User, error) {
	return model.User{
		UserId:       repository.UserId,
		EmailAddress: repository.EmailAddress,
	}, nil
}

func (repository UserRepositoryMock) GetUserByIdAndPassword(userId model.UserId, password string) (model.User, error) {
	if userId != repository.UserId || password != repository.Password {
		return model.User{}, api_error.InvalidRequestError(errors.CustomError{Message: UserNotFoundError})
	}
	return model.User{Activated: repository.Activated}, nil
}

func (repository UserRepositoryMock) CheckIfUserIsActivated(userId model.UserId) error {
	if repository.Activated {
		return api_error.InvalidRequestError(errors.CustomError{Message: UserIsAlreadyActivatedError})
	}
	return nil
}
