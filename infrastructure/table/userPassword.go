package table

type Password string

type UserPassword struct {
	UserId    UserId   `json:"user_id" gorm:"primary_key"`
	Password  Password `json:"password"`
	UpdatedAt string   `json:"updated_at" sql:"DEFAULT:current_timestamp"`
	CreatedAt string   `json:"created_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}
