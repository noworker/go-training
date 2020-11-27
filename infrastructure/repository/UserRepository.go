package repository

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib"
	"go_training/lib/errors"
	"go_training/web/api_error"
	"time"
)

const (
	CanNotCreateExistingUserId errors.ErrorMessage = "can_not_create_existing_user_id"
)

const activationTokenLifeTime = time.Hour

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) infrainterface.IUserRepository {
	return userRepository{
		DB: DB,
	}
}

func (repository userRepository) Activate(userId model.UserId) error {
	conn := map[string]interface{}{
		"user_id": userId,
	}
	result := repository.DB.Model(&table.User{}).Where(conn).Update("activated", true)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (repository userRepository) UserExists(userId model.UserId, password lib.HashedByteString) bool {
	userPassword := table.UserPassword{}
	conn := map[string]interface{}{
		"user_id":  userId,
		"password": password,
	}
	result := repository.DB.Where(conn).Find(&userPassword)
	if result.RecordNotFound() {
		return false
	}
	return true
}

func (repository userRepository) GetUserByIdAndPassword(userId model.UserId, password lib.HashedByteString) (table.User, error) {
	user := table.User{}
	conn := map[string]interface{}{
		"user_id":  userId,
		"password": password,
	}
	result := repository.DB.Where(conn).Find(&user)
	if err := result.Error; err != nil {
		return table.User{}, err
	}
	return user, nil
}

func (repository userRepository) CreateNewUser(user model.User, userPassword model.UserPassword) error {
	if exists := repository.UserExists(userPassword.UserId, userPassword.Password); exists {
		return api_error.InvalidRequestError(errors.CustomError{Message: CanNotCreateExistingUserId})
	}
	if err := repository.createUser(user.UserId, user.EmailAddress); err != nil {
		return api_error.InternalError(err)
	}
	if err := repository.createUserPassword(userPassword.UserId, userPassword.Password); err != nil {
		return api_error.InternalError(err)
	}
	return nil
}

func (repository userRepository) createUser(userId model.UserId, emailAddress model.EmailAddress) error {
	user := table.User{
		UserId:       table.UserId(userId),
		EmailAddress: table.EmailAddress(emailAddress),
		Activated:    false,
	}
	result := repository.DB.Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (repository userRepository) createUserPassword(userId model.UserId, password lib.HashedByteString) error {
	user := table.UserPassword{UserId: table.UserId(userId), Password: table.Password(password)}
	result := repository.DB.Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
