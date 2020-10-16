package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/web/api_error"
	"net/http"
)

type CreateUserHandler struct {
	UserRepository infrainterface.IUserRepository
}

type User struct {
	UserId       string `json:"user_id"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

func (handler CreateUserHandler) CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return api_error.InvalidRequestError(err)
	}

	newUser, err := model.NewUser(user.UserId, user.EmailAddress, user.Password)
	if err != nil {
		return api_error.InvalidRequestError(err)
	}
	if err := handler.UserRepository.CreateUnactivatedNewUser(newUser); err != nil {
		return api_error.InvalidRequestError(err)
	}

	return c.String(http.StatusCreated, "User is successfully created.")
}
