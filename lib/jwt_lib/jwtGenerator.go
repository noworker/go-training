package jwt_lib

import (
	"github.com/dgrijalva/jwt-go"
	mdate "github.com/matsuri-tech/date-go"
	"io/ioutil"
)

const ThisDir = "/home/ryo/matsuri/go-training/"

func Generator(userId string) (string, error) {
	signBytes, err := ioutil.ReadFile(ThisDir + "private.pem")
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     mdate.GetToday().PlusNDay(1).Unix(),
	})

	tokenString, err := token.SignedString(signBytes)
	if err != nil {
		println("hoge3")
		return "", err
	}

	return tokenString, nil
}
