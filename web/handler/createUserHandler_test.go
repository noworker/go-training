package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go_training/domain/model"
	"go_training/infrastructure/repository"
	"go_training/initializer"
	"go_training/lib"
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

	hashedPassword, err := lib.MakeHashedStringFromPassword(password)
	if err != nil {
		t.Error(err.Error())
	}

	newUser, err := model.NewUser(userId, emailAddress, hashedPassword)

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
	form := make(url.Values)

	userId := "abcde"
	emailAddress := "abc@example.com"

	form.Set("user_id", userId)
	form.Set("email_address", emailAddress)
	form.Set("password", "111")

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}

	c := e.NewContext(req, rec)
	err := handlers.CreateUserHandler.CreateUser(c)

	if err.(*echo.HTTPError).Message.(string) != string(lib.PasswordIsTooShort) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", lib.PasswordIsTooShort)
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

	c := e.NewContext(req, rec)
	err := handlers.CreateUserHandler.CreateUser(c)

	if err.(*echo.HTTPError).Message.(string) != string(model.InvalidEmailAddressFormat) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", model.InvalidEmailAddressFormat)
	}
}

func TestCreateUserHandlerErrorCase3(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)

	userId := "abcde"

	form := make(url.Values)
	form.Set("user_id", userId)
	form.Set("email_address", "aaa@example.com")

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}

	c := e.NewContext(req, rec)
	err := handlers.CreateUserHandler.CreateUser(c)

	if err.(*echo.HTTPError).Message.(string) != string(lib.PasswordIsTooShort) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", lib.PasswordIsTooShort)
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

	c := e.NewContext(req, rec)
	err := handlers.CreateUserHandler.CreateUser(c)

	if err.(*echo.HTTPError).Message.(string) != string(repository.CanNotCreateExistingUserId) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", repository.CanNotCreateExistingUserId)
	}
}

func TestCreateUserHandlerErrorCase5(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)
	userId := "aaa"
	emailAddress := "abc@example.com"
	password := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	form := make(url.Values)
	form.Set("user_id", userId)
	form.Set("email_address", emailAddress)
	form.Set("password", password)

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}

	c := e.NewContext(req, rec)
	err := handlers.CreateUserHandler.CreateUser(c)

	if err.(*echo.HTTPError).Message.(string) != string(lib.PasswordItTooLong) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", lib.PasswordItTooLong)
	}
}

func TestCreateUserHandlerErrorCase6(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)

	userId := "12345678901234567"
	emailAddress := "abc@example.com"
	password := "12345678"

	form := make(url.Values)
	form.Set("user_id", userId)
	form.Set("email_address", emailAddress)
	form.Set("password", password)

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}

	c := e.NewContext(req, rec)
	err := handlers.CreateUserHandler.CreateUser(c)

	if err.(*echo.HTTPError).Message.(string) != string(model.UserIdIsTooLong) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", model.UserIdIsTooLong)
	}
}
