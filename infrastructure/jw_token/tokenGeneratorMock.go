package jw_token

type TokenGeneratorMock struct {
}

func NewTokenGeneratorMock(path string) (TokenGeneratorMock, error) {
	return TokenGeneratorMock{}, nil
}

func (g TokenGeneratorMock) GenerateActivateUserToken(userId string) (string, error) {
	return "token", nil
}
