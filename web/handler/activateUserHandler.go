package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/service"
	"go_training/web/api_error"
	"net/http"
)

type ActivateUserHandler struct {
	activateUserService service.ActivateUserService
}

func (handler ActivateUserHandler) ActivateUser(c echo.Context) error {
	token := c.QueryParam("token")

	err := handler.activateUserService.ActivateUser(token)
	if err != nil {
		return api_error.InternalError(err)
	}

	return c.String(http.StatusCreated, "User is successfully Activated.")
}
