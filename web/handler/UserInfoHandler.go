package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/lib/errors"
	"go_training/web/api_error"
	"net/http"
)

const (
	UserIsNotActivated errors.ErrorMessage = "user_is_not_activated"
	TokenIsInvalid     errors.ErrorMessage = "token_is_invalid"
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
	userId := c.Param("user_id")
	token, err := c.Cookie("login_token")
	if err != nil {
		return api_error.InvalidRequestError(errors.CustomError{
			Message: TokenIsInvalid,
			Option:  err.Error(),
		})
	}
	if _, err := handler.tokenChecker.CheckLoginUserToken(model.Token(token.Value)); err != nil {
		return err
	}

	user, err := handler.userRepository.GetUserById(model.UserId(userId))
	if err != nil {
		return err
	}

	if !user.Activated {
		return api_error.InvalidRequestError(errors.CustomError{
			Message: UserIsNotActivated,
		})
	}

	return c.JSON(http.StatusOK, UserInfo{
		UserId:       string(user.UserId),
		EmailAddress: string(user.EmailAddress),
	})
}
