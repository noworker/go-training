package handler

import "go_training/initializer"

type Handlers struct {
	CreateUserHandler CreateUserHandler
}

func InitHandler(repositories initializer.Repositories) Handlers {
	createUserRepository := repositories.UserRepository
	createUserHandler := CreateUserHandler{
		UserRepository: createUserRepository,
	}

	handlers := Handlers{
		CreateUserHandler: createUserHandler,
	}
	return handlers
}
