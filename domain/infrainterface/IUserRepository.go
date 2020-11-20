package infrainterface

import (
	"go_training/domain/model"
)

type IUserRepository interface {
	Activate(userId model.UserId) error
	//CheckIfActivated(userId model.UserId, password lib.HashStringPassword) (bool, error)
	//GetUserByIdAndPassword(userId model.UserId, password lib.HashStringPassword) (table.User, error)
	CreateNewUser(user model.User) error
}
