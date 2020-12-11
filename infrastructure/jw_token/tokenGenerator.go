package jw_token

import (
	"github.com/dgrijalva/jwt-go"
	"go_training/domain/model"
	"go_training/lib/errors"
	"io/ioutil"
	"strconv"
)

const (
	NoFileError          errors.ErrorMessage = "no file"
	ParsePrivateKeyError errors.ErrorMessage = "parse private privateKey error"
	EncodeTokenError     errors.ErrorMessage = "encode token error"
)

type TokenType string

const (
	Activation   TokenType = "Activation"
	Verification TokenType = "Verification"
	Login        TokenType = "Login"
)

type TokenGenerator struct {
	privateKeyPath []byte
}

func NewTokenGenerator(path string) (TokenGenerator, error) {
	signBytes, err := ioutil.ReadFile(path + "private.pem")
	if err != nil {
		return TokenGenerator{}, errors.CustomError{Message: NoFileError, Option: err.Error()}
	}
	return TokenGenerator{privateKeyPath: signBytes}, nil
}

func (g TokenGenerator) generateToken(userId model.UserId, tokenType TokenType, expiresAt model.UnixTime) (model.Token, error) {
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(g.privateKeyPath)
	if err != nil {
		return "", errors.CustomError{Message: ParsePrivateKeyError, Option: err.Error()}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"type":    tokenType,
		"exp":     strconv.Itoa(int(expiresAt)),
	})

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", errors.CustomError{Message: EncodeTokenError, Option: err.Error()}
	}

	return model.Token(tokenString), nil
}

func (g TokenGenerator) GenerateActivateUserToken(userId model.UserId) (model.Token, error) {
	return g.generateToken(userId, Activation, model.CurrentUnixTime().AddMinutes(10))
}

func (g TokenGenerator) GenerateLoginUserToken(userId model.UserId) (model.Token, error) {
	return g.generateToken(userId, Login, model.CurrentUnixTime().AddHours(24*30))
}

func (g TokenGenerator) GenerateTwoStepVerificationToken(userId model.UserId) (model.Token, error) {
	return g.generateToken(userId, Verification, model.CurrentUnixTime().AddMinutes(10))
}
