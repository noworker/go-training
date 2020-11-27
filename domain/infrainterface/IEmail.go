package infrainterface

type IEmail interface {
	SendEmail(to, token string)
}
