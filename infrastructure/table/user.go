package table

type UserId string
type EmailAddress string
type Password string

type User struct {
	UserId       UserId
	EmailAddress EmailAddress
	Password     Password
	Activated    bool
}