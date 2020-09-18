package model

type Password HashedString

type UserPassword struct {
	UserId
	Password
}
