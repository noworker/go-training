package infrainterface

type IEmail interface {
	SendEmail(to, tokenURL string)
}
