package infrastructure

import (
	"net/smtp"
)

type mail struct {
	from     string
	username string
	password string
	to       string
	sub      string
	msg      string
}

type EMailSender struct {
	emailAddress string
	password     string
}

func NewEmailSender(address, password string) EMailSender {
	return EMailSender{emailAddress: address, password: password}
}

func (sender EMailSender) SendEmail(to string) error {
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth(sender.emailAddress, sender.emailAddress, sender.password, "smtp.gmail.com")

	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail("smtp.gmail.com:587", auth, sender.emailAddress, []string{to}, MessageBuilder(to))
	if err != nil {
		return err
	}
	return nil
}

func MessageBuilder(to string) []byte {
	return []byte("" +
		"From: <" + "API実装テスト" + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 勉強のために自動メール送信を行っています。もし身に覚えがない場合、削除してください。すみません。\r\n" +
		"内容" + "\r\n")
}
