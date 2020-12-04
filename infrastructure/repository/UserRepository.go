package repository

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib"
	"go_training/lib/errors"
	"go_training/web/api_error"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	CanNotCreateExistingUserId errors.ErrorMessage = "can_not_create_existing_user_id"
	UserNotFoundError          errors.ErrorMessage = "user not found"
	InvalidPassword            errors.ErrorMessage = "invalid password"
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

func (repository userRepository) UserExists(userId model.UserId, password string) bool {
	userPassword := table.UserPassword{}
	conn := map[string]interface{}{
		"user_id": userId,
	}
	result := repository.DB.Where(conn).Find(&userPassword)
	if result.RecordNotFound() {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (repository userRepository) getUserPassword(userId model.UserId, password string) error {
	userPassword := table.UserPassword{}
	conn := map[string]interface{}{
		"user_id": userId,
	}

	result := repository.DB.Where(conn).Find(&userPassword)
	if err := result.Error; err != nil {
		return errors.CustomError{Message: UserNotFoundError,
			Option: "UserRepository:66"}
	}

	if result.RecordNotFound() {
		return errors.CustomError{Message: UserNotFoundError}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPassword.Password), []byte(password)); err != nil {
		return errors.CustomError{Message: InvalidPassword}
	}

	return nil
}

func (repository userRepository) GetUserByIdAndPassword(userId model.UserId, password string) (model.User, error) {
	if err := repository.getUserPassword(userId, password); err != nil {
		return model.User{}, err
	}

	user := table.User{}
	conn := map[string]interface{}{
		"user_id": userId,
	}

	result := repository.DB.Where(conn).First(&user)
	if err := result.Error; err != nil {
		return model.User{}, errors.CustomError{Message: UserNotFoundError,
			Option: "UserRepository:80"}
	}

	if result.RecordNotFound() {
		return model.User{}, errors.CustomError{Message: UserNotFoundError}
	}

	return user.MapToModel(), nil
}

func (repository userRepository) CreateNewUser(user model.User, rawPassword string, hashedPassword lib.HashedByteString) error {
	if exists := repository.UserExists(user.UserId, rawPassword); exists {
		return api_error.InvalidRequestError(errors.CustomError{Message: CanNotCreateExistingUserId})
	}
	if err := repository.createUser(user.UserId, user.EmailAddress); err != nil {
		return api_error.InternalError(err)
	}
	if err := repository.createUserPassword(user.UserId, hashedPassword); err != nil {
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
