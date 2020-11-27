package email

import (
	"go_training/domain/infrainterface"
	"go_training/domain/model"
)

type SenderMock struct {
}

func NewEmailSenderMock(address model.EmailAddress, password string) infrainterface.IEmail {
	return SenderMock{}
}

func (sender SenderMock) SendEmail(address model.EmailAddress, token model.Token) {
}
