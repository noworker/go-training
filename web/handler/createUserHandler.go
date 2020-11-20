package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go_training/config"
	"go_training/domain/service"
	"go_training/lib/jwt_lib"
	"go_training/web/api_error"
	"net/http"
)

type CreateUserHandler struct {
	conf              config.Config
	createUserService service.CreateUserService
}

type User struct {
	UserId       string `json:"user_id" form:"user_id"`
	EmailAddress string `json:"email_address" form:"email_address"`
	Password     string `json:"password" form:"password"`
}

func (handler CreateUserHandler) CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return api_error.InvalidRequestError(err)
	}
	if err := handler.createUserService.CreateUser(user.UserId, user.EmailAddress, user.Password); err != nil {
		return api_error.InvalidRequestError(err)
	}
	token, err := jwt_lib.Generator(user.UserId, handler.conf)
	if err != nil {
		return api_error.InvalidRequestError(err)
	}
	return c.String(http.StatusCreated, fmt.Sprintf("User is successfully created.\ntoken: %s", token))
}
