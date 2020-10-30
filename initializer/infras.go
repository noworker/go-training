package initializer

import (
	"go_training/config"
	"go_training/infrastructure"
)

type Infras struct {
	EmailSender infrastructure.EMailSender
}

func InitInfras(conf *config.Config) Infras {
	emailSender := infrastructure.NewEmailSender(conf.App.SenderEmailAddress, conf.App.EmailPassword)
	return Infras{
		EmailSender: emailSender,
	}
}
