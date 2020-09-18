package repository

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/infrastructure/table"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) infrainterface.IUserRepository {
	return  UserRepository{
		DB: DB,
	}
}

func (repository UserRepository) Activate(userId model.UserId, password model.Password) error {
	var user table.User{}
	conn := map[string]interface{} {
		"user_id": userId,
		"password": password,
		"activated": true,
	}
	// アカウントの存在を予め確認したほうが良いかも……
	result := repository.DB.Where(conn).Save(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (repository UserRepository) CheckIfActivated(userId model.UserId, password model.Password) (bool, error) {
	user, err := repository.GetUserByIdAndPassword(userId, password)
	if err != nil {
		return false, err
	}
	return user.Activated, nil
}

func (repository UserRepository) GetUserByIdAndPassword(userId model.UserId, password model.Password) (table.User, error) {
	var user table.User{}
	conn := map[string]interface{} {
		"user_id": userId,
		"password": password,
	}
	result := repository.DB.Where(conn).Find(&user)
	if err := result.Error; err != nil {
		return table.User{}, err
	}
	return user, nil
}

func (repository UserRepository) SaveUnactivatedNewUser(userId model.UserId, emailAddress model.EmailAddress, password model.Password) error {
	var user table.User{}
	conn := map[string]interface{} {
		"user_id": userId,
		"email_address": emailAddress,
		"password": password,
		"activated": false,
	}
	result := repository.DB.Where(conn).Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}