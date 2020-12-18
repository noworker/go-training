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

func (sender SenderMock) SendActivationEmail(address model.EmailAddress, token model.Token) {
}

func (sender SenderMock) SendTwoStepVerificationEmail(address model.EmailAddress, token model.Token) {
}
