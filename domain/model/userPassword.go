package model

type HashString string

type UserPassword struct {
	UserId
	Password HashString
}
