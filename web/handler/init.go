package handler

import (
	"go_training/initializer"
)

type Handlers struct {
	CreateUserHandler   CreateUserHandler
	ActivateUserHandler ActivateUserHandler
}

func InitHandler(repositories initializer.Repositories, services initializer.Services, infras initializer.Infras) Handlers {
	createUserHandler := CreateUserHandler{
		createUserService: services.CreateUserService,
	}

	activateUserHandler := ActivateUserHandler{
		activateUserService: services.ActivateUserService,
	}

	handlers := Handlers{
		CreateUserHandler:   createUserHandler,
		ActivateUserHandler: activateUserHandler,
	}
	return handlers
}
