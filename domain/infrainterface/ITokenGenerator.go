package infrainterface

import "go_training/domain/model"

type ITokenGenerator interface {
	GenerateActivateUserToken(userId model.UserId) (model.Token, error)
}
