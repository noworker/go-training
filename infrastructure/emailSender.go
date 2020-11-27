package infrastructure

import (
	"net/smtp"
)

type EMailSender struct {
	emailAddress string
	password     string
}

func NewEmailSender(address, password string) EMailSender {
	return EMailSender{emailAddress: address, password: password}
}

func (sender EMailSender) SendEmail(to, tokenURL string) error {
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
