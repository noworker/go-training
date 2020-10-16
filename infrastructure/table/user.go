package table

type UserId string
type EmailAddress string

type User struct {
	UserId       UserId       `json:"user_id" gorm:"primary_key"`
	EmailAddress EmailAddress `json:"email_address"`
	Activated    bool         `json:"activated"`
	UpdatedAt    string       `json:"updated_at" sql:"DEFAULT:current_timestamp"`
	CreatedAt    string       `json:"created_at" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}
