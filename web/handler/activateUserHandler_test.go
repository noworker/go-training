package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go_training/domain/model"
	"go_training/domain/service"
	"go_training/infrastructure/jw_token"
	"go_training/infrastructure/repository"
	"go_training/initializer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func initActivateUserHandlerMock(userId, token string) Handlers {
	tokenChecker, _ := jw_token.NewTokenCheckerMock(model.UserId(userId), token)
	repo := repository.NewUserRepositoryMock(userId, "password", "aaa@example.com")
	activateUserService := service.NewActivateUserService(tokenChecker, repo)
	return InitHandler(
		initializer.Repositories{},
		initializer.Services{
			ActivateUserService: activateUserService,
		},
		initializer.Infras{})
}

func TestActivateUserHandler_ActivateUser(t *testing.T) {
	token := "token"
	handlers := initActivateUserHandlerMock("user_id", token)
	e := NewRouter(handlers)
	q := make(url.Values)

	q.Set("token", token)

	req := httptest.NewRequest(http.MethodGet, "/api/activate_user"+"?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusOK, rec.Code) {
		t.Error(rec)
	}
}
