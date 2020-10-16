package repository

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib/errors"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) infrainterface.IUserRepository {
	return  userRepository{
		DB: DB,
	}
}

func (repository userRepository) Activate(userId model.UserId, password model.HashString) error {
	if exists, err := repository.userExists(userId, password); exists {
		return err
	}
	user := table.User{
		Activated: true,
	}
	conn := map[string]interface{} {
		"user_id": userId,
	}
	result := repository.DB.Where(conn).Save(&user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (repository userRepository) userExists(userId model.UserId, password model.HashString) (bool, error) {
	userPassword := table.UserPassword{}
	conn := map[string]interface{} {
		"user_id": userId,
		"password": password,
	}
	result := repository.DB.Where(conn).Find(&userPassword)
	if result.RecordNotFound() {
		return false, nil
	}
	return true, errors.CustomError{Message: "can not create existing user id", ErrorType: errors.CanNotCreateExistingUserId}
}

//func (repository userRepository) CheckIfActivated(userId model.UserId, password model.HashStringPassword) (bool, error) {
//	user, err := repository.GetUserByIdAndPassword(userId, password)
//	if err != nil {
//		return false, err
//	}
//	return user.Activated, nil
//}

//func (repository userRepository) GetUserByIdAndPassword(userId model.UserId, password model.HashStringPassword) (table.User, error) {
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

func (repository userRepository) CreateUnactivatedNewUser(userId model.UserId, emailAddress model.EmailAddress, password model.HashString) error {
	if exists, err := repository.userExists(userId, password); exists {
		return err
	}
	if err := repository.createUser(userId, emailAddress); err != nil {
		return err
	}
	if err := repository.createUserPassword(userId, password); err != nil {
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
		panic(err.Error())
		return err
	}
	return nil
}

func (repository userRepository) createUserPassword(userId model.UserId, password model.HashString) error {
	user := table.UserPassword{UserId: table.UserId(userId), Password: table.Password(password)}
	result := repository.DB.Create(&user)
	if err := result.Error; err != nil {
		panic(err.Error())
		return err
	}
	return nil
}