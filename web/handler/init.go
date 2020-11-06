package handler

import "go_training/initializer"

type Handlers struct {
	CreateUserHandler CreateUserHandler
}

func InitHandler(repositories initializer.Repositories, services initializer.Services) Handlers {
	createUserHandler := CreateUserHandler{
		createUserService: services.CreateUserService,
	}

	handlers := Handlers{
		CreateUserHandler: createUserHandler,
	}
	return handlers
}
