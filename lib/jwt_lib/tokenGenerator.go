package jwt_lib

import (
	"github.com/dgrijalva/jwt-go"
	mdate "github.com/matsuri-tech/date-go"
	"go_training/config"
	"io/ioutil"
)

func Generator(userId string, conf config.Config) (string, error) {
	signBytes, err := ioutil.ReadFile(conf.App.KeyPath + "private.pem")
	if err != nil {
		return "", err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     mdate.GetToday().PlusNDay(1).Unix(),
	})

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
