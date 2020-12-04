package infrainterface

import "go_training/domain/model"

type ITokenChecker interface {
	CheckActivateUserToken(jwtStr model.Token) (model.UserId, error)
	CheckLoginUserToken(jwtStr model.Token) (model.UserId, error)
}
