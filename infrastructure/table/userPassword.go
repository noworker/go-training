package table

type Password string

type UserPassword struct {
	UserId       UserId 		`json:"user_id"`
	Password     Password		`json:"password"`
}