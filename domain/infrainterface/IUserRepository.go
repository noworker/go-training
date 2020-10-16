package infrainterface

import (
	"go_training/domain/model"
)

type IUserRepository interface {
	Activate(userId model.UserId, password model.HashString) error
	//CheckIfActivated(userId model.UserId, password model.HashStringPassword) (bool, error)
	//GetUserByIdAndPassword(userId model.UserId, password model.HashStringPassword) (table.User, error)
	CreateUnactivatedNewUser(user model.User) error
}
