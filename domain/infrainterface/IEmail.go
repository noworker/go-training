package infrainterface

import "go_training/domain/model"

type IEmail interface {
	SendActivationEmail(address model.EmailAddress, token model.Token)
	SendTwoStepVerificationEmail(address model.EmailAddress, token model.Token)
}
