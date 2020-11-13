package lib

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	mdate "github.com/matsuri-tech/date-go"
	"go_training/domain/model"
	"io/ioutil"
)

type JWT string

func JWTGenerator(userId model.UserId) (JWT, error) {
	signBytes, err := ioutil.ReadFile("./jwt_key.rsa")
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userId,
		"activate": true,
		"nbf":      mdate.GetToday().Unix(),
	})

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return JWT(tokenString), nil
}
