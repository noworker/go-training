package handler

import (
	"go_training/domain/model"
	"go_training/domain/service"
	"go_training/infrastructure/jw_token"
	"go_training/infrastructure/repository"
	"go_training/initializer"
)

func initResendActivationEmailHandlerMock(tokenCheckerUserId, repositoryUserId, token string, activated bool) Handlers {
	tokenChecker, _ := jw_token.NewTokenCheckerMock(model.UserId(tokenCheckerUserId), token)
	repo := repository.NewUserRepositoryMock(repositoryUserId, "password", "aaa@example.com", activated)
	activateUserService := service.NewActivateUserService(tokenChecker, repo)
	return InitHandler(
		initializer.Repositories{},
		initializer.Services{
			ActivateUserService: activateUserService,
		},
		initializer.Infras{})
}
