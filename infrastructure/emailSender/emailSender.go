package email

import (
	"github.com/labstack/gommon/log"
	"go_training/domain/infrainterface"
	"go_training/domain/model"
	"net/smtp"
)

type Sender struct {
	emailAddress string
	password     string
}

func NewEmailSender(address, password string) infrainterface.IEmail {
	return Sender{emailAddress: address, password: password}
}

func (sender Sender) SendEmail(address model.EmailAddress, token model.Token) {
	tokenURL := "http://localhost:8080/activate_user?token=" + token
	auth := smtp.PlainAuth(sender.emailAddress, sender.emailAddress, sender.password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender.emailAddress, []string{string(address)}, messageBuilder(address, tokenURL))
	if err != nil {
		log.Printf("failed to send email %v", err.Error())
	}
}

func messageBuilder(address model.EmailAddress, tokenURL model.Token) []byte {
	return []byte("" +
		"From: <" + "API実装テスト" + ">\r\n" +
		"To: " + string(address) + "\r\n" +
		"Subject: 勉強のために自動メール送信を行っています。もし身に覚えがない場合、削除してください。すみません。\r\n" +
		"内容" + "\r\n" +
		string(tokenURL) + "\r\n")
}
