package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go_training/infrastructure/repository"
	"go_training/initializer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const existingUserId = "existing"

func TestCreateUserHandler(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)
	form := make(url.Values)
	form.Set("user_id", "Saburo")
	form.Set("email_address", "saburo@example.com")
	form.Set("password", "12345678")
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusCreated, rec.Code){
		t.Error("error")
	}
	if !assert.JSONEq(t,
		`{"user_id": "Saburo", "email_address": "saburo@example.com", "password": "12345678"}`,
		rec.Body.String()) {
		t.Error("error")
	}
}
