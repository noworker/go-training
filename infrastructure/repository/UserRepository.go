package repository

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib"
	"go_training/lib/errors"
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

func (repository userRepository) Activate(userId model.UserId, password lib.HashedByteString) error {
	if exists, err := repository.userExists(userId, password); !exists {
		return err
	}
	user := table.User{
		Activated: true,
	}
	conn := map[string]interface{}{
		"user_id": userId,
	}
	result := repository.DB.Where(conn).Save(&user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (repository userRepository) createEmailActivationToken(userId model.UserId, token lib.Token) error {
	EmailActivation := table.EmailActivationToken{
		ActivationToken: table.ActivationToken(token),
		UserId:          table.UserId(userId),
		ExpiresAt:       time.Now().Add(activationTokenLifeTime).Unix(),
	}
	result := repository.DB.Create(&EmailActivation)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (repository userRepository) userExists(userId model.UserId, password lib.HashedByteString) (bool, error) {
	userPassword := table.UserPassword{}
	conn := map[string]interface{}{
		"user_id":  userId,
		"password": password,
	}
	result := repository.DB.Where(conn).Find(&userPassword)
	if result.RecordNotFound() {
		return false, nil
	}
	return true, errors.CustomError{Message: CanNotCreateExistingUserId}
}

//func (repository userRepository) CheckIfActivated(userId model.UserId, password lib.HashStringPassword) (bool, error) {
//	user, err := repository.GetUserByIdAndPassword(userId, password)
//	if err != nil {
//		return false, err
//	}
//	return user.Activated, nil
//}

//func (repository userRepository) GetUserByIdAndPassword(userId model.UserId, password lib.HashStringPassword) (table.User, error) {
//	user := table.User{}
//	conn := map[string]interface{} {
//		"user_id": userId,
//		"password": password,
//	}
//	result := repository.DB.Where(conn).Find(&user)
//	if err := result.Error; err != nil {
//		return table.User{}, err
//	}
//	return user, nil
//}

func (repository userRepository) CreateUnactivatedNewUser(user model.User, userPassword model.UserPassword, token lib.Token) error {
	if exists, err := repository.userExists(userPassword.UserId, userPassword.Password); exists {
		return err
	}
	if err := repository.createUser(user.UserId, user.EmailAddress); err != nil {
		return err
	}
	if err := repository.createUserPassword(userPassword.UserId, userPassword.Password); err != nil {
		return err
	}
	if err := repository.createEmailActivationToken(user.UserId, token); err != nil {
		return err
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
