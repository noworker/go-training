package infrainterface

import (
	"go_training/domain/model"
)

type IUserRepository interface {
	Activate(userId model.UserId, password model.Password) error
	//CheckIfActivated(userId model.UserId, password model.Password) (bool, error)
	//GetUserByIdAndPassword(userId model.UserId, password model.Password) (table.User, error)
	CreateUnactivatedNewUser(userId model.UserId, emailAddress model.EmailAddress, password model.Password) error
}
