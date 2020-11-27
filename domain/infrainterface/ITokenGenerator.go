package infrainterface

type ITokenGenerator interface {
	GenerateActivateUserToken(userId string) (string, error)
}
