package handler

import (
	"github.com/labstack/echo/v4"
	"go_training/domain/model"
	"go_training/domain/service"
	"net/http"
)

type VerificationHandler struct {
	service.TwoStepVerificationService
}
type VerifyResponse struct {
	LoginToken model.Token `json:"login_token"`
}

func (handler VerificationHandler) Verify(c echo.Context) error {
	token := c.QueryParam("token")
	loginToken, err := handler.TwoStepVerificationService.Verify(model.Token(token))
	if err != nil {
		return err
	}

	response := VerifyResponse{LoginToken: loginToken}

	return c.JSON(http.StatusOK, response)
}
