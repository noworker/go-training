package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/web/api_error"
	"net/http"
)

type UserToken struct {
	UserId string `json:"user_id" form:"user_id"`
	Token  string `json:"toke" form:"token"`
}

type UserInfo struct {
	UserId       string `json:"user_id"`
	EmailAddress string `json:"email_address"`
}

type UserInfoHandler struct {
	userRepository infrainterface.IUserRepository
	tokenChecker   infrainterface.ITokenChecker
}

func (handler UserInfoHandler) GetUserInfo(c echo.Context) error {
	userToken := new(UserToken)
	if err := c.Bind(userToken); err != nil {
		return api_error.InvalidRequestError(err)
	}

	if err := handler.userRepository.CheckIfUserIsActivated(model.UserId(userToken.UserId)); err != nil {
		return err
	}

	if _, err := handler.tokenChecker.CheckLoginUserToken(model.Token(userToken.Token)); err != nil {
		return err
	}

	user, err := handler.userRepository.GetUserById(model.UserId(userToken.UserId))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, UserInfo{
		UserId:       string(user.UserId),
		EmailAddress: string(user.EmailAddress),
	})
}
