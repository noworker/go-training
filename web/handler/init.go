package handler

import (
	"go_training/config"
	"go_training/initializer"
)

type Handlers struct {
	CreateUserHandler   CreateUserHandler
	ActivateUserHandler ActivateUserHandler
}

func InitHandler(repositories initializer.Repositories, services initializer.Services, conf config.Config) Handlers {

	createUserHandler := CreateUserHandler{
		conf:              conf,
		createUserService: services.CreateUserService,
	}

	activateUserHandler := ActivateUserHandler{
		conf:                 conf,
		createUserRepository: repositories.UserRepository,
	}

	handlers := Handlers{
		CreateUserHandler:   createUserHandler,
		ActivateUserHandler: activateUserHandler,
	}
	return handlers
}
