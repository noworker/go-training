package infrainterface

type ITokenChecker interface {
	CheckActivateUserToken(jwtStr string) (string, error)
}
