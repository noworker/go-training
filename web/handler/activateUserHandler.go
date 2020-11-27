package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/config"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/infrastructure/jw_token"
	"go_training/web/api_error"
	"net/http"
)

type ActivateUserHandler struct {
	conf                 config.Config
	createUserRepository infrainterface.IUserRepository
}

func (handler ActivateUserHandler) ActivateUser(c echo.Context) error {
	token := c.QueryParam("token")
	userId, err := jw_token.CheckActivateUserToken(token, handler.conf)
	if err != nil {
		return api_error.InvalidRequestError(err)
	}
	if err := handler.createUserRepository.Activate(model.UserId(userId)); err != nil {
		return api_error.InvalidRequestError(err)
	}
	return c.String(http.StatusCreated, "User is successfully Activated.")
}
