package initializer

import (
	"go_training/config"
	"go_training/domain/infrainterface"
	email "go_training/infrastructure/emailSender"
	"go_training/infrastructure/jw_token"
)

type Infras struct {
	EmailSender    infrainterface.IEmail
	TokenGenerator infrainterface.ITokenGenerator
	TokenChecker   infrainterface.ITokenChecker
}

func InitInfras(conf config.Config) Infras {
	emailSender := email.NewEmailSender(conf.App.SenderEmailAddress, conf.App.EmailPassword)
	tokenGenerator, err := jw_token.NewTokenGenerator(conf.App.KeyPath)
	if err != nil {
		panic(err.Error())
	}

	tokenChecker, err := jw_token.NewTokenChecker(conf.App.KeyPath)
	if err != nil {
		panic(err.Error())
	}

	return Infras{
		EmailSender:    emailSender,
		TokenGenerator: tokenGenerator,
		TokenChecker:   tokenChecker,
	}
}
