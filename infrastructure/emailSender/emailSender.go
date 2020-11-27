package email

import (
	"go_training/domain/infrainterface"
	"net/smtp"
)

type Sender struct {
	emailAddress string
	password     string
}

func NewEmailSender(address, password string) infrainterface.IEmail {
	return Sender{emailAddress: address, password: password}
}

func (sender Sender) SendEmail(to, token string) error {
	tokenURL := "http://localhost:8080/activate_user?token=" + token
	auth := smtp.PlainAuth(sender.emailAddress, sender.emailAddress, sender.password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender.emailAddress, []string{to}, messageBuilder(to, tokenURL))
	if err != nil {
		return err
	}
	return nil
}

func messageBuilder(to, tokenURL string) []byte {
	return []byte("" +
		"From: <" + "API実装テスト" + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 勉強のために自動メール送信を行っています。もし身に覚えがない場合、削除してください。すみません。\r\n" +
		"内容" + "\r\n" +
		tokenURL + "\r\n")
}
