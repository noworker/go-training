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

func initActivateUserHandlerMock(tokenCheckerUserId, repositoryUserId, token string, activated bool) Handlers {
	tokenChecker, _ := jw_token.NewTokenCheckerMock(model.UserId(tokenCheckerUserId), token)
	repo := repository.NewUserRepositoryMock(repositoryUserId, "password", "aaa@example.com", activated)
	activateUserService := service.NewActivateUserService(tokenChecker, repo)
	return InitHandler(
		initializer.Repositories{},
		initializer.Services{
			ActivateUserService: activateUserService,
		},
		initializer.Infras{})
}

func ActivateUserHandlerTester(repoMockUserId, tokenMockUserId, token, mockToken string) httptest.ResponseRecorder {
	handlers := initActivateUserHandlerMock(repoMockUserId, tokenMockUserId, mockToken, false)
	e := NewRouter(handlers)
	q := make(url.Values)

	q.Set("token", token)

	req := httptest.NewRequest(http.MethodGet, "/api/activate_user"+"?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	return *rec
}

func TestActivateUserHandler_ActivateUser(t *testing.T) {
	rec := ActivateUserHandlerTester("userId", "userId", "token", "token")
	if !assert.Equal(t, http.StatusOK, rec.Code) {
		t.Error(rec)
	}

	rec = ActivateUserHandlerTester("userId", "hoge", "token", "token")
	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error(rec)
	}

	rec = ActivateUserHandlerTester("userId", "userId", "token", "hoge")
	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error(rec)
	}
}
