package infrainterface

import (
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib"
)

type IUserRepository interface {
	Activate(userId model.UserId) error
	UserExists(userId model.UserId, password lib.HashedByteString) bool
	GetUserByIdAndPassword(userId model.UserId, password lib.HashedByteString) (table.User, error)
	CreateNewUser(user model.User, userPassword model.UserPassword) error
}
