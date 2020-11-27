package infrainterface

import "go_training/domain/model"

type IEmail interface {
	SendEmail(address model.EmailAddress, token model.Token)
}
