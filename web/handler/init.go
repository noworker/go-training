package handler

import (
	"go_training/initializer"
)

type Handlers struct {
	CreateUserHandler            CreateUserHandler
	ActivateUserHandler          ActivateUserHandler
	ResendActivationEmailHandler ResendActivationEmailHandler
	LoginHandler                 LoginHandler
}

func InitHandler(repositories initializer.Repositories, services initializer.Services, infras initializer.Infras) Handlers {
	createUserHandler := CreateUserHandler{
		createUserService: services.CreateUserService,
	}

	activateUserHandler := ActivateUserHandler{
		activateUserService: services.ActivateUserService,
	}

	resendActivationEmailHandler := ResendActivationEmailHandler{
		resendActivationEmailService: services.ResendActivationEmailService,
	}

	loginHandler := LoginHandler{
		userRepository: repositories.UserRepository,
		tokenGenerator: infras.TokenGenerator,
	}

	handlers := Handlers{
		CreateUserHandler:            createUserHandler,
		ActivateUserHandler:          activateUserHandler,
		ResendActivationEmailHandler: resendActivationEmailHandler,
		LoginHandler:                 loginHandler,
	}
	return handlers
}
