package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go_training/domain/service"
	email "go_training/infrastructure/emailSender"
	"go_training/infrastructure/jw_token"
	"go_training/infrastructure/repository"
	"go_training/initializer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func initResendActivationEmailHandlerMock(userId string, activated bool) Handlers {
	repo := repository.NewUserRepositoryMock(userId, "password", "aaa@example.com", activated)
	tokenGenerator, _ := jw_token.NewTokenGeneratorMock("aaa")
	emailSender := email.NewEmailSenderMock("aaa@example.com", "password")
	resendService := service.ResendActivationEmailService{
		UserRepository: repo,
		TokenGenerator: tokenGenerator,
		EmailSender:    emailSender,
	}
	return InitHandler(
		initializer.Repositories{},
		initializer.Services{
			ResendActivationEmailService: resendService,
		},
		initializer.Infras{})
}

func ResendActivationEmailHandlerTester(mockUserId, userId string, activated bool) *httptest.ResponseRecorder {
	handlers := initResendActivationEmailHandlerMock(mockUserId, activated)
	e := NewRouter(handlers)
	form := make(url.Values)

	form.Set("user_id", userId)
	form.Set("email_address", "aaa@example.com")
	form.Set("password", "password")

	req := httptest.NewRequest(http.MethodPost, "/api/resend_activation_email", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	return rec
}

func TestResendActivationEmailHandler_ResendActivationEmail(t *testing.T) {
	rec := ResendActivationEmailHandlerTester("user_id", "user_id", false)

	if !assert.Equal(t, http.StatusOK, rec.Code) {
		t.Error("error")
	}

	rec = ResendActivationEmailHandlerTester("user_id", "hoge", false)
	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}

	rec = ResendActivationEmailHandlerTester("user_id", "user_id", true)
	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}
