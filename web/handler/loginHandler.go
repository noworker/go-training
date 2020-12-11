package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/service"
	"go_training/web/api_error"
	"net/http"
)

type LoginHandler struct {
	loginService service.LoginService
}

type UserPassword struct {
	UserId   string `json:"user_id" form:"user_id"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	Status string `json:"status"`
}

func (handler LoginHandler) Login(c echo.Context) error {
	user := new(UserPassword)

	if err := c.Bind(user); err != nil {
		return api_error.InvalidRequestError(err)
	}

	if err := handler.loginService.Login(user.UserId, user.Password); err != nil {
		return err
	}

	response := LoginResponse{
		Status: "2段階認証Eメールが送信されました。",
	}

	return c.JSON(http.StatusOK, response)
}
