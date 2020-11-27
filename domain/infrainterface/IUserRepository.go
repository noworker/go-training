package infrainterface

import (
	"go_training/domain/model"
	"go_training/infrastructure/table"
	"go_training/lib"
)

type IUserRepository interface {
	Activate(userId model.UserId) error
	GetUserByIdAndPassword(userId model.UserId, password string) (table.User, error)
	CreateNewUser(user model.User, rawPassword string, hashedPassword lib.HashedByteString) error
}
