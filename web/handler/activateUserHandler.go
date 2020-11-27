package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/web/api_error"
	"net/http"
)

type ActivateUserHandler struct {
	tokenChecker         infrainterface.ITokenChecker
	createUserRepository infrainterface.IUserRepository
}

func (handler ActivateUserHandler) ActivateUser(c echo.Context) error {
	token := c.QueryParam("token")
	userId, err := handler.tokenChecker.CheckActivateUserToken(token)
	if err != nil {
		return api_error.InvalidRequestError(err)
	}
	if err := handler.createUserRepository.Activate(model.UserId(userId)); err != nil {
		return api_error.InvalidRequestError(err)
	}
	return c.String(http.StatusCreated, "User is successfully Activated.")
}
