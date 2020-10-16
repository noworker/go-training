package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go_training/domain/model"
	"go_training/infrastructure/repository"
	"go_training/initializer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const existingUserId = "existing"

func TestCreateUserHandlerNoErrorCase(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)
	form := make(url.Values)

	userId := "abcde"
	emailAddress := "abc@example.com"
	password := "12345678"

	form.Set("user_id", userId)
	form.Set("email_address", emailAddress)
	form.Set("password", password)

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusCreated, rec.Code) {
		t.Error("error")
	}

	newUser, err := model.NewUser(userId, emailAddress, password)

	if err != nil {
		t.Error("error")
		return
	}

	if repo.User.UserId != newUser.UserId {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.User.UserId, newUser.UserId))
	}

	if repo.User.EmailAddress != newUser.EmailAddress {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.User.EmailAddress, newUser.EmailAddress))
	}

	if repo.User.Password != newUser.Password {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.User.Password, newUser.Password))
	}
}

func TestCreateUserHandlerErrorCase1(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)
	form1 := make(url.Values)

	userId := "abcde"
	emailAddress := "abc@example.com"

	form1.Set("user_id", userId)
	form1.Set("email_address", emailAddress)
	form1.Set("password", "111")

	req1 := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form1.Encode()))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec1 := httptest.NewRecorder()

	e.ServeHTTP(rec1, req1)

	if !assert.Equal(t, http.StatusBadRequest, rec1.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase2(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)

	userId := "abcde"
	password := "12345678"

	form := make(url.Values)
	form.Set("user_id", userId)
	form.Set("email_address", "aaaexample.com")
	form.Set("password", password)

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase3(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)

	userId := "abcde"

	form := make(url.Values)
	form.Set("user_id", userId)
	form.Set("email_address", "aaaexample.com")

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase4(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)

	emailAddress := "abc@example.com"
	password := "12345678"

	form := make(url.Values)
	form.Set("user_id", existingUserId)
	form.Set("email_address", emailAddress)
	form.Set("password", password)

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}
