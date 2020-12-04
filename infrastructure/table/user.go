package table

import "go_training/domain/model"

type UserId string
type EmailAddress string

type User struct {
	UserId       UserId       `json:"user_id" gorm:"primary_key"`
	EmailAddress EmailAddress `json:"email_address"`
	Activated    bool         `json:"activated"`
	UpdatedAt    string       `json:"updated_at" sql:"DEFAULT:current_timestamp"`
	CreatedAt    string       `json:"created_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}

func (user User) MapToModel() model.User {
	return model.User{
		UserId:       model.UserId(user.UserId),
		EmailAddress: model.EmailAddress(user.EmailAddress),
		Activated:    user.Activated,
	}
}
