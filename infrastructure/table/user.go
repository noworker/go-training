package table

type UserId string
type EmailAddress string

type User struct {
	UserId       UserId 		`json:"user_id"`
	EmailAddress EmailAddress	`json:"email_address"`
	Activated    bool			`json:"activated"`
	UpdatedAt	 string			`json:"updated_at" sql:"not null;type:date"`
	CreatedAt	 string			`json:"created_at" sql:"not null;type:date"`
}