package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib"
	"go_training/web/api_error"
	"net/http"
)

type CreateUserHandler struct {
	UserRepository infrainterface.IUserRepository
}

type User struct {
	UserId  string `json:"user_id" form:"user_id" validate:"required, user_id"`
	EmailAddress string `json:"email_address" form:"email_address" validate:"required, email_address"`
	Password string `json:"password" form:"password" validate:"required, password"`
}

func(handler CreateUserHandler) CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return api_error.InvalidRequestError(err)
	}

	if err := c.Validate(user); err != nil {
		return api_error.InvalidRequestError(err)
	}

	if err := handler.UserRepository.CreateUnactivatedNewUser(
		model.UserId(user.UserId),
		model.EmailAddress(user.EmailAddress),
		lib.MakeHashedStringFromPassword(user.Password),
		); err != nil {
		return api_error.InvalidRequestError(err)
	}
	
	return c.JSON(http.StatusOK, user)
}