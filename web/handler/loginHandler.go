package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/web/api_error"
	"net/http"
)

type LoginHandler struct {
	userRepository infrainterface.IUserRepository
	tokenGenerator infrainterface.ITokenGenerator
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

	if _, err := handler.userRepository.GetUserByIdAndPassword(model.UserId(user.UserId), user.Password); err != nil {
		return err
	}

	token, err := handler.tokenGenerator.GenerateLoginUserToken(model.UserId(user.UserId))
	if err != nil {
		return err
	}

	response := LoginResponse{
		UserId: user.UserId,
		Token:  string(token),
	}

	return c.JSON(http.StatusOK, response)
}
