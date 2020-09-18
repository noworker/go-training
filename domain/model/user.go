package model

type UserId string
type Password HashedString
type EmailAddress string

type User struct {
	UserId
	Password
	EmailAddress
}