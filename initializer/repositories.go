package initializer

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/infrastructure/repository"
)

type Repositories struct {
	UserRepository infrainterface.IUserRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	userRepository := repository.NewUserRepository(db)
	return Repositories{
		UserRepository: userRepository,
	}
}
