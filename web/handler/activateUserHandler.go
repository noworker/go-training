package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/model"
	"go_training/domain/service"
	"net/http"
)

type ActivateUserHandler struct {
	activateUserService service.ActivateUserService
}

func (handler ActivateUserHandler) ActivateUser(c echo.Context) error {
	token := c.QueryParam("token")
	err := handler.activateUserService.ActivateUser(model.Token(token))
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "User is successfully Activated.")
}
