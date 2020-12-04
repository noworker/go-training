package infrainterface

import (
	"go_training/domain/model"
	"go_training/lib"
)

type IUserRepository interface {
	Activate(userId model.UserId) error
	GetUserByIdAndPassword(userId model.UserId, password string) (model.User, error)
	CreateNewUser(user model.User, rawPassword string, hashedPassword lib.HashedByteString) error
	CheckIfUserIsActivated(userId model.UserId) error
}
