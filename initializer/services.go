package initializer

import (
	"go_training/domain/service"
)

type Services struct {
	CreateUserService   service.CreateUserService
	ActivateUserService service.ActivateUserService
}

func InitServices(repositories Repositories, infras Infras) Services {
	createUserService := service.CreateUserService{
		UserRepository: repositories.UserRepository,
		TokenGenerator: infras.TokenGenerator,
		EmailSender:    infras.EmailSender,
	}
	activateUserService := service.ActivateUserService{
		UserRepository: repositories.UserRepository,
		TokenChecker:   infras.TokenChecker,
	}
	return Services{
		CreateUserService:   createUserService,
		ActivateUserService: activateUserService,
	}
}
