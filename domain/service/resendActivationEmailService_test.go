package service

import (
	"go_training/domain/model"
	email "go_training/infrastructure/emailSender"
	"go_training/infrastructure/jw_token"
	"go_training/infrastructure/repository"
	"testing"
)

func initResendService(userId, password, address string) ResendActivationEmailService {
	userRepository := repository.NewUserRepositoryMock(userId, password, address)
	tokenGenerator, err := jw_token.NewTokenGeneratorMock("path")
	if err != nil {
		panic(err)
	}
	emailSender := email.NewEmailSenderMock(model.EmailAddress(address), password)

	return ResendActivationEmailService{UserRepository: userRepository, TokenGenerator: tokenGenerator, EmailSender: emailSender}
}

func TestResendEmail(t *testing.T) {
	service := initResendService("user_id", "password", "address")
	err := service.ResendActivationEmail("user_id", "password", "address")
	if err != nil {
		t.Error(err)
	}
	err = service.ResendActivationEmail("user_i", "password", "address")
	if err == nil {
		t.Error("error")
	}
	err = service.ResendActivationEmail("user_id", "passwords", "address")
	if err == nil {
		t.Error("error")
	}
	err = service.ResendActivationEmail("user_i", "password", "addresss")
	if err == nil {
		t.Error("error")
	}

}
