package initializer

import (
	"go_training/domain/service"
)

type Services struct {
	CreateUserService service.CreateUserService
}

func InitServices(repositories Repositories) Services {
	createUserService := service.CreateUserService{
		UserRepository: repositories.UserRepository,
	}
	return Services{
		CreateUserService: createUserService,
	}
}
