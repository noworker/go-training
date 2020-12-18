package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go_training/domain/model"
	"go_training/domain/service"
	"net/http"
	"time"
)

type VerificationHandler struct {
	service.TwoStepVerificationService
}

func (handler VerificationHandler) Verify(c echo.Context) error {
	token := c.QueryParam("token")
	userId, loginToken, err := handler.TwoStepVerificationService.Verify(model.Token(token))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = string(loginToken)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/api/user_info?user_id=%s&token=%s", userId, loginToken))
}
