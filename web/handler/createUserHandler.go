package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib"
	"net/http"
)

type CreateUserHandler struct {
	UserRepository infrainterface.IUserRepository
}

type User struct {
	UserId  string `json:"user_id" form:"user_id"`
	EmailAddress string `json:"email_address" form:"email_address"`
	Password string `json:"password" form:"password"`
}

func(handler CreateUserHandler) CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	fmt.Printf("userId Is %s",user.UserId)
	err := handler.UserRepository.CreateUnactivatedNewUser(
		model.UserId(user.UserId),
		model.EmailAddress(user.EmailAddress),
		lib.MakeHashedString(user.Password),
		)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}