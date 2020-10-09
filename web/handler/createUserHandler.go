package handler

import (
	"go_training/domain/infrainterface"
	"github.com/labstack/echo/v4"
	"go_training/domain/model"
	"go_training/lib"
	"net/http"
)

type CreateUserHandler struct {
	UserRepository infrainterface.IUserRepository
}

func(handler CreateUserHandler) CreateUser(c echo.Context) error {
	userId := c.Param("user_id")
	password := c.Param("password")
	emailAddress:= c.Param("email_address")
	err := handler.UserRepository.CreateUnactivatedNewUser(
		model.UserId(userId),
		model.EmailAddress(emailAddress),
		lib.MakeHashedString(password),
		)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "200 OK")
}