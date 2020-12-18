package handler

import (
	"go_training/initializer"
)

type Handlers struct {
	CreateUserHandler
	ActivateUserHandler
	ResendActivationEmailHandler
	LoginHandler
	UserInfoHandler
	VerificationHandler
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
		loginService: services.LoginService,
	}

	userInfoHandler := UserInfoHandler{
		userRepository: repositories.UserRepository,
		tokenChecker:   infras.TokenChecker,
	}

	verificationHandler := VerificationHandler{
		TwoStepVerificationService: services.TwoStepVerificationService,
	}

	handlers := Handlers{
		CreateUserHandler:            createUserHandler,
		ActivateUserHandler:          activateUserHandler,
		ResendActivationEmailHandler: resendActivationEmailHandler,
		LoginHandler:                 loginHandler,
		UserInfoHandler:              userInfoHandler,
		VerificationHandler:          verificationHandler,
	}
	return handlers
}
