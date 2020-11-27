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

func (handler ResendActivationEmailHandler) ResentActivationEmail(c echo.Context) error {
	userId := c.QueryParam("user_id")
	password := c.QueryParam("password")
	emailAddress := c.QueryParam("email_address")

	err := handler.resendActivationEmailService.ResendActivationEmail(model.UserId(userId), password, model.EmailAddress(emailAddress))
	if err != nil {
		return api_error.InternalError(err)
	}

	return c.String(http.StatusCreated, "email is successfully sent.")
}
