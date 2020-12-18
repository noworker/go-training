package initializer

import (
	"go_training/domain/service"
)

type Services struct {
	service.CreateUserService
	service.ActivateUserService
	service.ResendActivationEmailService
	service.LoginService
	service.TwoStepVerificationService
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

	resendActivationEmailService := service.ResendActivationEmailService{
		UserRepository: repositories.UserRepository,
		TokenGenerator: infras.TokenGenerator,
		EmailSender:    infras.EmailSender,
	}

	loginService := service.LoginService{
		UserRepository: repositories.UserRepository,
		TokenGenerator: infras.TokenGenerator,
		EmailSender:    infras.EmailSender,
	}

	twoStepVerificationService := service.TwoStepVerificationService{
		TokenChecker:   infras.TokenChecker,
		TokenGenerator: infras.TokenGenerator,
	}
	return Services{
		CreateUserService:            createUserService,
		ActivateUserService:          activateUserService,
		ResendActivationEmailService: resendActivationEmailService,
		LoginService:                 loginService,
		TwoStepVerificationService:   twoStepVerificationService,
	}
}
