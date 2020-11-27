package handler

import (
	"go_training/config"
	"go_training/infrastructure/jw_token"
	"go_training/initializer"
)

type Handlers struct {
	CreateUserHandler   CreateUserHandler
	ActivateUserHandler ActivateUserHandler
}

func InitHandler(repositories initializer.Repositories, services initializer.Services, conf config.Config) Handlers {
	tokenGenerator, err := jw_token.NewTokenGenerator(conf.App.KeyPath)
	if err != nil {
		panic(err.Error())
	}

	tokenChecker, err := jw_token.NewTokenChecker(conf.App.KeyPath)
	if err != nil {
		panic(err.Error())
	}

	createUserHandler := CreateUserHandler{
		tokenGenerator:    tokenGenerator,
		createUserService: services.CreateUserService,
	}

	activateUserHandler := ActivateUserHandler{
		tokenChecker:         tokenChecker,
		createUserRepository: repositories.UserRepository,
	}

	handlers := Handlers{
		CreateUserHandler:   createUserHandler,
		ActivateUserHandler: activateUserHandler,
	}
	return handlers
}
