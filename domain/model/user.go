package model

type UserId string
type EmailAddress string

type User struct {
	UserId
	EmailAddress
	Activated bool
}