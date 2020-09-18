package initializer

import (
	"github.com/jinzhu/gorm"
	"go_training/domain/infrainterface"
	"go_training/infrastructure/repository"
)

type Repositories struct {
	userRepository infrainterface.IUserRepository
}

func InitRepositories(db *gorm.DB) Repositories {
	userRepository := repository.NewUserRepository(db)
	return Repositories{
		userRepository: userRepository,
	}
}