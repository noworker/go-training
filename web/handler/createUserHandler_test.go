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

	if !assert.Equal(t, http.StatusCreated, rec.Code){
		t.Error("error")
	}

	if repo.UserId != model.UserId(userId) {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.UserId, userId))
	}

	if repo.EmailAddress != model.EmailAddress(emailAddress) {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.EmailAddress, emailAddress))
	}

	if repo.Password != lib.MakeHashedStringFromPassword(password) {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.Password, password))
	}
}

func TestCreateUserHandlerErrorCase(t *testing.T) {
	repo := repository.NewUserRepositoryMock(existingUserId)
	handlers := InitHandler(initializer.Repositories{UserRepository: repo})
	e := NewRouter(handlers)
	form1 := make(url.Values)

	userId := "abcde"
	emailAddress := "abc@example.com"
	password := "12345678"

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

	form2 := make(url.Values)
	form2.Set("user_id", userId)
	form2.Set("email_address", "aaaexample.com")
	form2.Set("password", password)

	req2 := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form1.Encode()))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec2 := httptest.NewRecorder()

	e.ServeHTTP(rec2, req2)

	if !assert.Equal(t, http.StatusBadRequest, rec2.Code) {
		t.Error("error")
	}
}
