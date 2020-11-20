package jwt_lib

import (
	"go_training/config"
	"testing"
)

func TestJWT(t *testing.T) {
	conf := config.NewDummyConfig()
	KeyGenerator(conf)
	token, err := TokenGenerator("user_id", conf)
	if err != nil {
		t.Error(err.Error())
	}
	userId, err := TokenChecker(token, conf)
	if err != nil {
		t.Error(err.Error())
	}
	if userId != "user_id" {
		t.Error("error")
	}
}
