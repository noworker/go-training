package email

import (
	"go_training/domain/infrainterface"
)

type SenderMock struct {
}

func NewEmailSenderMock(address, password string) infrainterface.IEmail {
	return SenderMock{}
}

func (sender SenderMock) SendEmail(to, token string) error {
	return nil
}
