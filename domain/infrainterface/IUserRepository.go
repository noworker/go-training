package infrainterface

import (
	"go_training/domain/model"
	"go_training/lib"
)

type IUserRepository interface {
	Activate(userId model.UserId, password lib.HashedByteString) error
	//CheckIfActivated(userId model.UserId, password lib.HashStringPassword) (bool, error)
	//GetUserByIdAndPassword(userId model.UserId, password lib.HashStringPassword) (table.User, error)
	CreateNewUser(user model.User, token lib.Token) error
}
