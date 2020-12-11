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

func (sender Sender) SendActivationEmail(address model.EmailAddress, token model.Token) {
	tokenURL := "http://localhost:8080/api/activate_user?token=" + token
	sender.sendEmail(address, string(tokenURL))
}

func (sender Sender) SendTwoStepVerificationEmail(address model.EmailAddress, token model.Token) {
	tokenURL := "http://localhost:8080/api/verification?token=" + token
	sender.sendEmail(address, string(tokenURL))
}

func (sender Sender) sendEmail(address model.EmailAddress, url string) {
	auth := smtp.PlainAuth(sender.emailAddress, sender.emailAddress, sender.password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender.emailAddress, []string{string(address)}, messageBuilder(address, url))
	if err != nil {
		log.Printf("failed to send email %v", err.Error())
	}
}

func messageBuilder(address model.EmailAddress, tokenURL string) []byte {
	return []byte("" +
		"From: <" + "API実装テスト" + ">\r\n" +
		"To: " + string(address) + "\r\n" +
		"Subject: 勉強のために自動メール送信を行っています。もし身に覚えがない場合、削除してください。すみません。\r\n" +
		"内容" + "\r\n" +
		string(tokenURL) + "\r\n")
}
