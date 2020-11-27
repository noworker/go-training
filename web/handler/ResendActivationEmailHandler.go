package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/model"
	"go_training/domain/service"
	"go_training/web/api_error"
	"net/http"
)

type ResendActivationEmailHandler struct {
	resendActivationEmailService service.ResendActivationEmailService
}

func NewResendActivationEmailHandler(resendActivationEmailService service.ResendActivationEmailService) ResendActivationEmailHandler {
	return ResendActivationEmailHandler{resendActivationEmailService: resendActivationEmailService}
}

func (handler ResendActivationEmailHandler) ResendActivationEmail(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return api_error.InvalidRequestError(err)
	}

	err := handler.resendActivationEmailService.ResendActivationEmail(model.UserId(user.UserId), user.Password, model.EmailAddress(user.EmailAddress))
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "email is successfully sent.")
}
