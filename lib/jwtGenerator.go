package lib

import (
	"github.com/dgrijalva/jwt-go"
	mdate "github.com/matsuri-tech/date-go"
	"io/ioutil"
)

type JWTStr string

func JWTGenerator(userId string) (JWTStr, error) {
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

	return JWTStr(tokenString), nil
}
