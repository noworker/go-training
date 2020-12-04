package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"go_training/domain/service"
	email "go_training/infrastructure/emailSender"
	"go_training/infrastructure/jw_token"
	"go_training/infrastructure/repository"
	"go_training/initializer"
	"go_training/lib"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

const existingUserId = "existing"

func initCreateHandlerMock(repo infrainterface.IUserRepository) Handlers {
	tokenGenerator, _ := jw_token.NewTokenGeneratorMock("aaa")
	emailSender := email.NewEmailSenderMock("hoge", "huga")
	createUserService := service.NewCreateUserService(repo, tokenGenerator, emailSender)
	return InitHandler(
		initializer.Repositories{
			UserRepository: repo,
		},
		initializer.Services{
			CreateUserService: createUserService,
		}, initializer.Infras{TokenGenerator: tokenGenerator})
}

func CreateUserHandlerTester(userId, emailAddress, password string, repo *repository.UserRepositoryMock) (httptest.ResponseRecorder, error) {
	handlers := initCreateHandlerMock(repo)
	e := NewRouter(handlers)
	form := make(url.Values)

	form.Set("user_id", userId)
	form.Set("email_address", emailAddress)
	form.Set("password", password)

	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	c := e.NewContext(req, rec)
	return *rec, handlers.CreateUserHandler.CreateUser(c)
}

func TestCreateUserHandlerNoErrorCase(t *testing.T) {
	userId := "abcde"
	password := "12345678"
	emailAddress := "abc@example.com"
	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, _ := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if !assert.Equal(t, http.StatusCreated, rec.Code) {
		t.Error("error")
	}

	hashedPassword, err := lib.MakeHashedStringFromPassword(password)
	if err != nil {
		t.Error(err.Error())
	}

	newUser, err := model.NewUser(userId, emailAddress)
	if err != nil {
		t.Error(err.Error())
		return
	}
	newUserPassword := model.NewUserPassword(newUser.UserId, hashedPassword)

	if repo.User.UserId != newUser.UserId {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.User.UserId, newUser.UserId))
	}

	if repo.User.EmailAddress != newUser.EmailAddress {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.User.EmailAddress, newUser.EmailAddress))
	}

	if reflect.DeepEqual(repo.UserPassword.Password, newUserPassword.Password) {
		t.Error(fmt.Sprintf("\nresult: %s\nexpected: %s", repo.UserPassword.Password, newUserPassword.Password))
	}
}

func TestCreateUserHandlerErrorCase1(t *testing.T) {
	userId := "abcde"
	password := "111"
	emailAddress := "abc@example.com"
	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, err := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if err.(*echo.HTTPError).Message.(string) != string(lib.PasswordIsTooShort) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", lib.PasswordIsTooShort)
	}

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase2(t *testing.T) {
	userId := "abcde"
	password := "111"
	emailAddress := "abcexample.com"
	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, err := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if err.(*echo.HTTPError).Message.(string) != string(model.InvalidEmailAddressFormat) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", model.InvalidEmailAddressFormat)
	}

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase3(t *testing.T) {
	userId := "abcde"
	password := ""
	emailAddress := "abc@example.com"
	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, err := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if err.(*echo.HTTPError).Message.(string) != string(lib.PasswordIsTooShort) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", lib.PasswordIsTooShort)
	}

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase4(t *testing.T) {
	userId := existingUserId
	password := "12345678"
	emailAddress := "abc@example.com"
	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, err := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if err.(*echo.HTTPError).Message.(string) != string(repository.CanNotCreateExistingUserId) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", repository.CanNotCreateExistingUserId)
	}

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase5(t *testing.T) {
	userId := "aaa"
	emailAddress := "abc@example.com"
	password := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"

	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, err := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if err.(*echo.HTTPError).Message.(string) != string(lib.PasswordIsTooLong) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", lib.PasswordIsTooLong)
	}

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}

func TestCreateUserHandlerErrorCase6(t *testing.T) {
	userId := "12345678901234567"
	emailAddress := "abc@example.com"
	password := "12345678"

	repo := repository.NewUserRepositoryMock(existingUserId, "aaa", "aaa", false)
	rec, err := CreateUserHandlerTester(userId, emailAddress, password, repo)

	if err.(*echo.HTTPError).Message.(string) != string(model.UserIdIsTooLong) {
		t.Error("\nresult", err.(*echo.HTTPError).Message, "\nexpect:", model.UserIdIsTooLong)
	}

	if !assert.Equal(t, http.StatusBadRequest, rec.Code) {
		t.Error("error")
	}
}
