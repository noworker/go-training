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
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func (handler LoginHandler) Login(c echo.Context) error {
	user := new(UserPassword)

	if err := c.Bind(user); err != nil {
		return api_error.InvalidRequestError(err)
	}

	token, err := handler.loginService.Login(user.UserId, user.Password)
	if err != nil {
		return err
	}

	response := LoginResponse{
		UserId: user.UserId,
		Token:  string(token),
	}

	return c.JSON(http.StatusOK, response)
}
